import React, { useState, useEffect } from 'react';
import logo from './logo.svg';
import { Classes, Icon, Intent, ITreeNode, Position, Tooltip, Tree, Button } from '@blueprintjs/core';
import { FolderPicker } from './FolderPicker';
import { MainNavBar } from './MainNavBar';
import { ipcRenderer } from 'electron';
import { SongGrid } from './SongGrid';
import { applyTheme } from '../themes/themes';
import { lightTheme } from '../themes/light';
import { darkTheme } from '../themes/dark';
import { isLight } from '../themes/colorMixer';
import { wrapGrid } from 'animate-css-grid';
import { Song } from '../models/song';

const themeName = 'dark';
export const theme = darkTheme;
applyTheme(themeName);

const App: React.FC<{}> = () => {
  const [selectedGrid, setSelectedGrid] = useState('song');
  const [themeDetails, setThemeDetails] = useState(isLight(theme.backgroundMain));
  const [sidePanelWidth, setSidePanelWidth] = useState(0);
  const [gridCols, setGridCols] = useState(`0px ${window.innerWidth}px`);
  const [gridClasses, setGridClasses] = useState('grid');
  const [songs, setSongs] = useState<Song[]>([]);
  const [gridMargin, setGridMargin] = useState(0);

  const gridRef = React.createRef<HTMLDivElement>();

  const updateTheme = (newThemeName: string) => {
    applyTheme(newThemeName);
    const newTheme = newThemeName === 'light' ? lightTheme : darkTheme;
    setThemeDetails(isLight(newTheme.backgroundMain));
  };
  // https://www.cssscript.com/pure-css-full-window-page-slider-with-folder-tab-navigation/

  useEffect(() => {
    if (gridRef.current) {
      const { unwrapGrid } = wrapGrid(gridRef.current, {
        duration: 150,
        onStart: () => {
          if (gridMargin > 0) {
            setGridMargin(0);
          }
        },
        onEnd: () => {
          if (gridMargin === 0) {
            setGridMargin(200);
          }
        },
      });
      // Remove animations after resizing because they don't play nicely with the virtualized grid
      setTimeout(unwrapGrid, 1);
    }
    setGridCols(`${sidePanelWidth}px ${window.innerWidth - sidePanelWidth}px`);
    if (sidePanelWidth > 0) {
      setGridClasses('expanded');
    } else {
      setGridClasses('collapsed');
    }
  }, [sidePanelWidth]);

  return (
    <>
      <MainNavBar
        sidePanelWidth={sidePanelWidth}
        setSidePanelWidth={setSidePanelWidth}
        selectedGrid={selectedGrid}
        setSelectedGrid={setSelectedGrid}
        updateTheme={updateTheme}
        isLight={themeDetails}
        songs={songs}
        setSongs={setSongs}
      />
      <div
        ref={gridRef}
        className={gridClasses}
        style={{
          paddingTop: 40,
          display: 'grid',
          gridTemplateRows: `${window.innerHeight - 110}px 70px`,
          gridTemplateColumns: gridCols,
        }}
      >
        <div>
          <div style={{ display: sidePanelWidth > 0 ? 'block' : 'none' }}>test</div>
        </div>
        <SongGrid
          selectedGrid={selectedGrid}
          isLightTheme={themeDetails}
          width={window.innerWidth - gridMargin}
          songs={songs}
          setSongs={setSongs}
        />
      </div>
    </>
  );
};

export default App;