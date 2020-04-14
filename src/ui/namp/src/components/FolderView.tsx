import React, { useState, useEffect } from 'react';
import { Table, Column, Cell } from '@blueprintjs/table';
import { Button, ITreeNode, Tooltip, Position, Icon, Classes, Intent, Toaster, Toast, ButtonGroup, Divider, Alert } from '@blueprintjs/core';
import { FolderPicker } from './FolderPicker';
import { getJson, putJson } from '../fetchUtil';
import { SelectedFolders } from './SelectedFolders';

const AppToaster = Toaster.create({
    position: Position.TOP
});


export const FolderView: React.FC<{width: number, height: number}> = ({width, height}) => {
    const [rows, setRows] = useState<Array<string>>([]);
    const [selected, setSelected] = useState<string>('');
    const [errorText, setErrorText] = useState<string>('');
    
    useEffect(() => {
      refreshFolders();
    }, []);

    const cellRenderer = (rowIndex: number) => {
        return <Cell>{rows[rowIndex]}</Cell>
    };

    const refreshFolders = () => getJson<Array<string>>('/configuredFolders').then(setRows);

    const addFolderClick = () => {
        rows.push(selected);
        setRows([...rows]);
    }

    const saveFoldersClick = async () => {
        try {
            await putJson<void>('/updateFolders', {folders: rows});
            AppToaster.show({message: 'Success', intent: Intent.SUCCESS, icon: 'tick-circle', timeout: 1000});
        }
        catch (e) {
            setErrorText(e.message);
        }
        
    }

    const revertClick = () => {
        refreshFolders();
    }
    const spacerWidth = 5;
    const panelWidth = (width - spacerWidth) / 2;
    const buttonHeight = 30;
    const buttonPanelHeight = 105;
    return (
        <>
        <div style={{display: 'flex', alignItems: 'top', alignSelf: 'center', width, height, marginTop: 20}}>
            <div style={{display: 'flex', flexDirection: 'column', width: panelWidth}}>
                <div style={{height: height-buttonPanelHeight}}>
                    <SelectedFolders rows={rows} setRows={setRows} width={panelWidth}/>
                </div>
                <div style={{display: 'flex', flexDirection: 'row', alignItems: 'center', flexGrow: 1}}>
                    <Button intent={Intent.SUCCESS} icon='floppy-disk' text='Save' style={{height: buttonHeight}}
                        onClick={saveFoldersClick}/>
                    <div style={{margin:spacerWidth}}/>
                    <Button intent={Intent.WARNING} icon='undo' text='Revert' style={{height: buttonHeight}} onClick={revertClick}/>
                </div>
            </div>
            <div style={{ width: spacerWidth}}/>
            <div style={{display: 'flex', flexDirection: 'column', width: panelWidth}}>
                <div style={{height: height-buttonPanelHeight}}>
                    <FolderPicker setSelected={setSelected}/>
                </div>
                <div style={{display: 'flex', flexDirection: 'row', alignItems: 'center', flexGrow: 1}}>
                    <Button intent={Intent.PRIMARY} onClick={addFolderClick} icon='add' text='Add' style={{height: buttonHeight}}/>
                </div>
            </div>
        </div>
        <Alert intent={Intent.DANGER} isOpen={errorText.length > 0} className={`bp3-dark`} onClose={() => setErrorText('')}>
            {errorText}
        </Alert>
    </>
    )
}
