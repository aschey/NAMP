mod models;
use dirs::home_dir;
use std::{sync::mpsc, str};
use actix_web::{dev::Server, HttpServer, HttpRequest, HttpResponse, App, http::{Method, header}, Result, http::StatusCode, get, error, web::Query, Responder, body, middleware};
use actix_cors::Cors;
use actix_files as fs;
use serde::{Deserialize, Serialize};
use std::fs::{read_dir, DirEntry, File, copy, remove_file};
use std::path::PathBuf;
use dotenv::dotenv;
use std::{vec::Vec, env};
use diesel;
use diesel::sqlite::SqliteConnection;
use diesel::prelude::*;
use models::{folder::*, mount::*, import::*};
use crate::schema::folder::dsl::*;
use crate::schema::mount::dsl::*;
use crate::schema::import_temp::dsl::*;
use crate::schema;
use sysinfo::{ProcessExt, SystemExt, DiskExt};
use itertools::Itertools;
use fstrings::*;
use std::io::prelude::*;
use actix_service::Service;
use futures::future::FutureExt;
use actix_http::http::header::{CacheControl, CacheDirective, HeaderName, HeaderValue};
use std::convert::TryFrom;
use async_std::prelude::*;
use futures::join;
use async_std::task;
use std::time::SystemTime;
use async_recursion::async_recursion;

use paperclip::actix::{
    // extension trait for actix_web::App and proc-macro attributes
    OpenApiExt, Apiv2Schema, api_v2_operation,
    // use this instead of actix_web::web
    web::{self, Json},
    api_v2_errors
};
use failure::Fail;

const IS_WINDOWS: bool = cfg!(windows);
const DATABASE_URL: &str = "DATABASE_URL";
fn get_delim() -> &'static str {
    return if IS_WINDOWS { "\\" } else { "/" };
}

fn convert_delim(path: String) -> String {
    return if IS_WINDOWS { path.replace("\\", "/") } else { path.replace("/", "\\") };
}

fn convert_delim_windows(path: String) -> String {
    return path.replace("/", "\\");
}

fn convert_delim_unix(path: String) -> String {
    return path.replace("\\", "/");
}

fn get_delim_escaped() -> &'static str {
    return if IS_WINDOWS { "\\\\" } else { "/" };
}

fn to_url_path(drive_path: String) -> String {
    if !IS_WINDOWS {
        return drive_path;
    }
    let replaced = convert_delim_unix(drive_path.replace(":", ""));
    return f!("/{replaced}");
}

#[cfg(windows)]
fn get_path() -> schema::folder::full_path_windows { full_path_windows }
#[cfg(unix)]
fn get_path() -> schema::folder::full_path_unix { full_path_unix }

#[cfg(windows)]
fn get_mount_path() -> schema::mount::windows_path { windows_path }
#[cfg(unix)]
fn get_mount_path() -> schema::mount::unix_path { unix_path }

pub fn establish_connection() -> SqliteConnection {
    dotenv().ok();
    let database_url = env::var(DATABASE_URL)
        .expect("DATABASE_URL must be set");
    SqliteConnection::establish(&database_url)
        .expect("Error connecting to database")

}

pub fn run_server(tx: mpsc::Sender<Server>) -> std::io::Result<()> {
    get_all_files();
    if !std::path::Path::new("./.env").exists() {
        write_env(&std::env::current_dir().unwrap().to_str().unwrap().to_owned());
    }

    let mut sys = actix_rt::System::new("server");

    let srv = HttpServer::new(|| { 
        let mut builder = App::new()
        .wrap(Cors::new().finish())
        .wrap_api()
        // REST endpoints
        .service(web::resource("/dirsInit").route(web::get().to(get_dirs_init)))
        .service(web::resource("/dirs").route(web::get().to(get_dirs)))
        .service(web::resource("/configuredFolders").route(web::get().to(get_configured_folders)))
        .service(web::resource("/isWindows").route(web::get().to(get_is_windows)))
        .service(web::resource("/updateFolders").route(web::put().to(update_folders)))
        .service(web::resource("/getDbPath").route(web::get().to(get_db_path)))
        .service(web::resource("/updateDbPath").route(web::put().to(update_db_path)))
        .service(web::resource("/getNtfsMounts").route(web::get().to(get_ntfs_mounts)))
        .service(web::resource("/updatePathMappings").route(web::put().to(update_path_mappings)))
        .with_json_spec_at("/spec")
        .build()
        // static files
        .wrap_fn(|req, srv| {
            let path = req.path().to_owned();
            srv.call(req).map(move |res| {
                if path == "/index.html" || path == "/" {
                    match res {
                        Ok(mut r) => {
                            r.headers_mut().insert(HeaderName::try_from("Cache-Control").unwrap(), HeaderValue::try_from("no-cache").unwrap());
                            Ok(r)
                        },
                        Err(r) => Err(r)
                    }
                    
                }
                else {
                    res
                }
            })
        })
        .service(fs::Files::new("/swagger", "./src/ui/namp/swagger").index_file("index.html"));

        let connection = establish_connection();
        let paths = folder.select(get_path()).load::<String>(&connection).unwrap();
        for path in paths {
            builder = builder.service(fs::Files::new(&to_url_path(path.to_owned()), path.to_owned()).show_files_listing());
        }
        let app = builder
            // Paths are matched in order so this needs to be last
            .service(fs::Files::new("/", "./src/ui/namp/build").index_file("index.html"));
        return app;
    })
    .bind("127.0.0.1:5000")?
    .run();

    // send server controller to main thread
    let _ = tx.send(srv.clone());

    // run future
    sys.block_on(srv)
}

fn get_timestamp(time: SystemTime) -> i32 {
    time.duration_since(SystemTime::UNIX_EPOCH).unwrap().as_secs() as i32
}

fn get_all_files_rec(start_path: String) -> Vec::<NewImport> {
    let mut all_files = Vec::<NewImport>::new();
    let mut dirs = std::fs::read_dir(start_path).unwrap();
    while let Some(dir_res) = dirs.next() {
        let dir = dir_res.unwrap();
        let path = dir.path();
        let full_path = path.to_str().unwrap();
        if path.is_file() {
            if full_path.ends_with(".mp3") || full_path.ends_with(".m4a") {
                let modified_time = path.metadata().unwrap().modified().unwrap();
                let created_time = dir.metadata().unwrap().created().unwrap();
                let last_changed = if modified_time > created_time { modified_time } else { created_time };
                let f = katatsuki::Track::from_path(std::path::Path::new(&path.to_str().unwrap()), None).unwrap();
                let n = NewImport {
                    import_artist: f.artist,
                    import_album: f.album,
                    import_modified_date: get_timestamp(last_changed),
                    import_song_path: to_url_path(full_path.to_owned()),
                    import_title: f.title
                };
                all_files.push(n);
            }
        }
        else {
            let inner_files = get_all_files_rec(full_path.to_owned());//.await;
            all_files = [all_files, inner_files].concat();
        }
    }
    return all_files;
}

async fn get_all_files_parallel(root: String) {
    //get_all_files_rec(root);
    //return;
    let proc_count = num_cpus::get() as f32;
    let dirs_and_files = read_dir(root).unwrap();

    let dirs = dirs_and_files.filter_map(|d| {
        let p = d.as_ref().unwrap().path();
        if p.is_dir() { Some(p.to_str().unwrap().to_owned()) } else { None }
    }).collect::<Vec<_>>();

    let dir_count = dirs.len() as f32;
    let dirs_per_thread = (dir_count / proc_count).ceil() as usize;
    let mut handles = Vec::<_>::new();
    println!("{}", dirs.len());
    for i in 0..proc_count as usize {
        let i = i.clone();
        let dirs = dirs.clone();
        handles.push(task::spawn(async move {
            let mut files = Vec::<_>::new();
            println!("{} {}", dirs_per_thread * i, dirs_per_thread*(i+1));
            for path in dirs[dirs_per_thread * i..std::cmp::min(dirs.len(), dirs_per_thread*(i+1))].to_vec() {
                files = [files, get_all_files_rec(path)].concat();
            }
            return files;
        }));
        //handles.push(get_all_files_rec2(dirs[dirs_per_thread * i..dirs_per_thread*(i+1)].to_vec()));
    }       
    let all = futures::future::join_all(handles).await;
    let res = all.iter().flatten().collect::<Vec<_>>();
    //return;
    let connection = establish_connection();
    let batch_size = 5000;
    let batches = (res.len() as f32 / batch_size as f32).ceil() as usize;
    for i in 0..batches {
        connection.transaction::<_, diesel::result::Error, _>(|| {
            //println!("{:?}", i * batch_size);
            //println!("{:?}", i * (batch_size + 1);
            let batch = res[i * batch_size..std::cmp::min(res.len(), (i + 1) * batch_size)].to_vec();
            for row in batch {
                let _r = diesel::insert_into(import_temp).values(row).execute(&connection).unwrap();
            }
            Ok(())
        }).unwrap();
    }
   
    
}

pub async fn get_all_files() {
    let now = std::time::Instant::now();
    let connection = establish_connection();
    let dirs = folder.select(get_path()).load::<String>(&connection).unwrap();
  
    get_all_files_parallel(dirs[0].to_owned()).await;
    // for dir in dirs {
    //     let all_files = get_all_files_rec(dir);
    //     for song in all_files {
    //         //let f = taglib::File::new(song.to_owned()).unwrap();
    //         let f = katatsuki::Track::from_path(std::path::Path::new(&song.to_owned()), None);
    //         // match f {
    //         //     Ok(t) => println!("{:?}", t.artist),
    //         //     Err(e) => println!("{:?}", e)
    //         // }
    //         // match tt {
    //         //     Ok(t) => println!("{:?}", t.artist()),
    //         //     Err(e) => println!("{:?}", e)
    //         // }
    //     }
    // }
    println!("{}", now.elapsed().as_secs());
}

fn filter_dirs(res: Result<DirEntry, std::io::Error>, delim: &str) -> Option<Dir> {
    let path = res.unwrap().path();
    let str_path = String::from(path.to_str().unwrap());
    let dir_name = String::from(str_path.split(delim).last().unwrap());
    if !dir_name.starts_with(".") { Some(Dir {name: dir_name, is_file: path.is_file() }) } else { None }
}

pub trait StrVecExt {
    fn sort_case_insensitive(&mut self);
}

impl StrVecExt for Vec<String> {
    fn sort_case_insensitive(&mut self) {
        &self.sort_by(|l, r| Ord::cmp(&l.to_lowercase(), &r.to_lowercase()));
    }
}

pub trait DirVecExt {
    fn sort_case_insensitive(&mut self);
}

impl DirVecExt for Vec<Dir> {
    fn sort_case_insensitive(&mut self) {
        &self.sort_by(|l, r| Ord::cmp(&l.name.to_lowercase(), &r.name.to_lowercase()));
    }
}

#[api_v2_errors(
    code=400,
    code=401, description="Unauthorized: Can't read session from header",
    code=500,
)]
#[derive(Fail, Debug)]
#[fail(display = "named error")]
struct HttpError {
    result: String,
}

// Use default implementation for `error_response()` method
impl error::ResponseError for HttpError {
    fn status_code(&self) -> StatusCode {
        StatusCode::BAD_REQUEST
    }

    fn error_response(&self) -> HttpResponse {
        let mut resp = HttpResponse::new(self.status_code());
        resp.headers_mut().insert(
            header::CONTENT_TYPE,
            header::HeaderValue::from_static("text/plain; charset=utf-8"),
        );
        resp.set_body(body::Body::from(self.result.to_owned()))
    }
}

fn get_dir_name(disk: &std::path::Path) -> String {
    let mut str_path = String::from(disk.to_str().unwrap());
    if IS_WINDOWS {
        str_path = str_path.replace("\\", "");
    }
    return str_path;
}

#[api_v2_operation]
async fn get_dirs_init() -> Result<Json<DirResponse>, ()> {
    let system = sysinfo::System::new_all();
    let disks = system.get_disks().iter().map(|d| Dir { is_file: false, name: get_dir_name(d.get_mount_point()) }).collect::<Vec<_>>();
    return Ok(Json(DirResponse {dirs: disks}))
}

#[api_v2_operation]
async fn get_dirs(dir_request: Query<DirRequest>) -> Result<Json<DirResponse>, ()> {
    let mut entries = read_dir(dir_request.dir.as_str()).unwrap()
        .filter_map(|res| filter_dirs(res, get_delim()))
        .collect::<Vec<_>>();

    entries.sort_case_insensitive();
    let response = Json(DirResponse {dirs: entries});
    return Ok(response);
}

#[api_v2_operation]
async fn get_is_windows() -> Result<Json<bool>, ()> {
    return Ok(Json(IS_WINDOWS));
}

fn get_configured_folders_helper() -> Vec<String> {
    let connection = establish_connection();
    let results = folder.load::<Folder>(&connection).expect("error");
    let paths = results.iter()
        .map(|rr| get_platform_folder(rr).clone())
        .filter(|r| r.len() > 0)
        .collect();
    return paths;
}

#[api_v2_operation]
async fn get_configured_folders() -> Result<Json<Vec<String>>, ()> {
    let paths = get_configured_folders_helper();
    return Ok(Json(paths));
}

fn get_ntfs_mounts_helper() -> Vec<String> {
    let system = sysinfo::System::new_all();
    let disks = system.get_disks();
    let fuse_disks = disks.iter()
        .filter(|d| str::from_utf8(d.get_file_system()).unwrap() == "fuseblk")
        .map(|d| get_dir_name(d.get_mount_point()))
        .collect::<Vec<_>>();
    let configured = get_configured_folders_helper();
    let configured_fuse = fuse_disks.into_iter()
        .filter(|f| configured.iter().any(|c| c.starts_with(f)))
        //.map(|f| NtfsMapping { dir: f, drive: "C:".to_owned()})
        .collect::<Vec<_>>();
    return configured_fuse;
}

#[api_v2_operation]
async fn get_ntfs_mounts() -> Result<Json<Vec<NtfsMapping>>, ()> {
    let connection = establish_connection();
    let mut fs_fuse = get_ntfs_mounts_helper();
    //fs_fuse.push("/mnt/test".to_owned());
    let mapped = mount.select((unix_path, windows_path)).load::<(String, String)>(&connection).unwrap();
    if IS_WINDOWS {
        let all = mapped.iter().map(|m| NtfsMapping { dir: m.0.to_owned(), drive: m.1.to_owned()}).collect::<Vec<_>>();
        return Ok(Json(all));
    }
    let mapped_unix = mapped.iter().map(|m| m.0.to_owned()).collect::<Vec<_>>();
    let mut mappings = fs_fuse.iter()
        .filter(|f| mapped_unix.contains(f))
        .map(|f| NtfsMapping { dir: f.to_owned(), drive: mapped.iter().filter(|m| m.0 == f.to_owned()).map(|m| m.1.to_owned()).collect()}).collect::<Vec<_>>();
    mappings.extend(
        fs_fuse.iter()
        .filter(|f| !mapped_unix.contains(f))
        .map(|f| NtfsMapping { dir: f.to_owned(), drive: "".to_owned()}).collect::<Vec<_>>()
    );
    return Ok(Json(mappings));
}

fn get_subfolders(new_folders: Vec<String>) -> Vec<String> {
    let copy = new_folders.to_vec();
    let dedup = &new_folders.into_iter().dedup_by(|l, r| r.starts_with(l)).collect::<Vec<_>>();
    
    let lala = copy.into_iter().filter(|f| !dedup.contains(f)).collect::<Vec<_>>();
    return lala;
}

fn get_dupe_folders(new_folders: Vec<String>) -> Vec<(String, Vec<String>)> {
    let grouped = new_folders.into_iter().group_by(|f| String::from(f)).into_iter().map(|(key, group)| (key, group.collect::<Vec<_>>())).collect::<Vec<(String, Vec<String>)>>();
    return grouped;
}

fn get_platform_folder(f: &Folder) -> String {
    if IS_WINDOWS { f.full_path_windows.to_owned() } else { f.full_path_unix.to_owned() }
}

fn new_folder(path: String) -> NewFolder {
    if IS_WINDOWS {
        NewFolder {
            full_path_unix: "".to_owned(),
            full_path_windows: path
        }
    }
    else {
        NewFolder {
            full_path_unix: path,
            full_path_windows: "".to_owned()
        }
    }
}

#[api_v2_operation]
async fn update_folders(new_folders_req: Json<FolderUpdate>) -> Result<Json<()>, HttpError> {
    let mut new_folders = new_folders_req.folders.to_vec();
    new_folders.sort_case_insensitive();
    let new_folders3 = new_folders.to_vec();
    let grouped = get_dupe_folders(new_folders);
    for (_, group) in grouped.into_iter() {
        if group.len() > 1 {
            let dup = group[0].to_owned();
            return Err(HttpError {result: f!("Duplicate folder chosen: {dup}")});
        }
    }

    let invalid_folders = get_subfolders(new_folders3);
    if invalid_folders.len() > 0 {
        let invalid = invalid_folders[0].to_owned();
        return Err(HttpError {result: f!("Unable to select a folder that is a child of another selected folder: {invalid}")});
    }

    let connection = establish_connection();
    //let sql = diesel::debug_query::<diesel::sqlite::Sqlite, _>(&folder.filter(get_path().ne_all(new_folders_req.folders.iter()).and(get_path().to_owned().ne("")))).to_string();
    //println!("{:?}", sql);
    let pred = folder.filter(
        get_path().ne_all(new_folders_req.folders.iter())
        .and(
            get_path().ne("")
        ));
    let to_remove = pred.to_owned().select(get_path()).load::<String>(&connection).unwrap();
    let all_mounts = mount.select(get_mount_path()).load::<String>(&connection).unwrap();
    for r in to_remove {
        let remove = all_mounts.iter().filter(|m| r.starts_with(m.to_owned())).collect::<Vec<_>>();
        let _ = diesel::delete(mount.filter(get_mount_path().eq_any(remove))).execute(&connection);
    }
    
    let res = diesel::delete(pred.to_owned()).execute(&connection);
    if res.is_err() {
        return Err(HttpError {result: "fail".to_owned()});
    }
    let existing = folder
        .filter(get_path().eq_any(new_folders_req.folders.iter()))
        .load::<Folder>(&connection).expect("error");
        
    let existing_paths = existing.iter().map(|rr| get_platform_folder(rr).clone()).collect::<Vec<_>>();
    let folders_to_create = new_folders_req.folders.iter()
        .filter(|f| !existing_paths.contains(f))
        .map(|f| new_folder(f.to_owned())).collect::<Vec<_>>();
    let res1 = diesel::insert_into(folder).values(folders_to_create).execute(&connection);
    if res1.is_err() {
        return Err(HttpError {result: "fail".to_owned()});
    }
    let r = mount
        .select((unix_path, windows_path))
        .load::<(String, String)>(&connection)
        .unwrap()
        .iter()
        .map(|f| NtfsMapping { dir: f.0.to_owned(), drive: f.1.to_owned()})
        .collect();
    sync_folder_mappings(r);
    return Ok(Json(()));
}

fn write_env(dir: &String) -> String {
    let mut file = File::create(".env").unwrap();
    let delim_escaped = get_delim_escaped();
    let escaped = dir.replace(get_delim(), get_delim_escaped());
    let full_url = f!("{escaped}{delim_escaped}namp.db");
    let _ = file.write_all(f!("DATABASE_URL={full_url}").as_bytes());
    return full_url;
}

#[api_v2_operation]
async fn update_db_path(request: Json<DirRequest>) -> Result<Json<()>, ()> {
    let full_url = write_env(&request.dir);
    let current_url = env::var(DATABASE_URL).unwrap();
    let _res = copy(current_url.to_owned(), full_url.to_owned());
    let _res2 = remove_file(current_url.to_owned());
    env::set_var(DATABASE_URL, full_url);
    return Ok(Json(()));
}

fn sync_folder_mappings(mapping: Vec<NtfsMapping>) {
    let connection = establish_connection();
    for r in mapping {
        if !IS_WINDOWS {
            let paths = folder
            .filter(full_path_unix.like(r.dir.to_owned() + "%"))
            .select((folder_id, full_path_unix))
            .load::<(i32, String)>(&connection).unwrap();
            for path in paths {
                let mut replace_val = "".to_owned();
                if r.drive != "" {
                    let suffix = convert_delim_windows(path.1.replace(&r.dir, ""));
                    replace_val = f!("{r.drive}{suffix}");
                }
                let _ = diesel::update(folder.filter(folder_id.eq(path.0))).set(full_path_windows.eq(replace_val)).execute(&connection);
            }
        }
        else {
            if r.drive == "" {
                continue;
            }
            let paths2 = folder
                .filter(full_path_windows.like(r.drive.to_owned() + "%"))
                .select((folder_id, full_path_windows))
                .load::<(i32, String)>(&connection).unwrap();
            for path in paths2 {
                let suffix = convert_delim_unix(path.1.replace(&r.drive, ""));
                let replace_val = f!("{r.dir}{suffix}");
                let _ = diesel::update(folder.filter(folder_id.eq(path.0))).set(full_path_unix.eq(replace_val)).execute(&connection);
            }

            let _ = diesel::update(folder.filter(full_path_unix.like(r.dir.to_owned() + "%").and(full_path_windows.not_like(r.drive.to_owned() + "%")))).set(full_path_unix.eq("")).execute(&connection);
        }
        
        
        // let _ = diesel::update(
        //     folder.filter(full_path_windows.like(f!("{r.drive}%"))
        //         .and(full_path_unix.not_like(f!("{r.dir}%")))))
        //         .set(full_path_unix.eq(f!"{r.dir}/"))
        //         .execute(&connection).unwrap();

    }
}

#[api_v2_operation]
async fn update_path_mappings(request: Json<Vec<NtfsMapping>>) -> Result<Json<()>, HttpError>  {
    let connection = establish_connection();
    let ins = request.iter().map(|r| (unix_path.eq(r.dir.to_owned()), windows_path.eq(r.drive.to_owned()))).collect::<Vec<_>>();
    let res = diesel::replace_into(mount).values(ins).execute(&connection);
    if res.is_err() {
        return Err(HttpError {result: "fail".to_owned()});
    }
    let pred = mount.filter(unix_path.ne_all(request.iter().map(|r| r.dir.to_owned())));
    if IS_WINDOWS {
        let to_delete = pred.to_owned().select(unix_path).load::<String>(&connection).unwrap();
        for d in to_delete {
            let _ = diesel::update(folder.filter(full_path_unix.like(d + "%"))).set(full_path_unix.eq("")).execute(&connection);
        }
        
    }
    
    let res2 = diesel::delete(pred).execute(&connection);
    if res2.is_err() {
        return Err(HttpError {result: "fail".to_owned()});
    }

    sync_folder_mappings(request.into_inner());
    return Ok(Json(()));
}

#[api_v2_operation]
async fn get_db_path() -> Result<Json<Dir>, ()>{
    let mut file = File::open(".env").unwrap();
    let mut contents = String::new();
    let _ = file.read_to_string(&mut contents);
    let delim_escaped = get_delim_escaped();
    let delim = get_delim();
    let res = contents.split("=").last().unwrap()
        .replace(delim_escaped, delim)
        .replace(&f!("{delim}namp.db"), "");
    return Ok(Json(Dir { is_file: true, name: res.to_owned()}));
}

#[derive(Serialize, Apiv2Schema)]
#[serde(rename_all = "camelCase")]
struct Dir {
    is_file: bool,
    name: String
}

#[derive(Serialize, Apiv2Schema)]
#[serde(rename_all = "camelCase")]
struct DirResponse {
    dirs: Vec<Dir>,
}

#[derive(Serialize, Deserialize, Apiv2Schema)]
#[serde(rename_all = "camelCase")]
struct NtfsMapping {
    dir: String,
    drive: String
}

#[derive(Deserialize, Apiv2Schema)]
#[serde(rename_all = "camelCase")]
struct DirRequest {
    dir: String,
}

#[derive(Deserialize, Apiv2Schema)]
#[serde(rename_all = "camelCase")]
struct FolderUpdate {
    folders: Vec<String>,
}