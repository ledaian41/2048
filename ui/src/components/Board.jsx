import React from 'react';
import {AnimatePresence} from 'framer-motion';
import {Tile} from './Tile';

export const Board = ({tiles, size}) => {
  const gridCells = Array.from({length: tiles.length}).map((_, i) => (
    <div key={`cell-${i}`} className="grid-cell"/>
  ));

  return (
    <div className="game-grid">
      <AnimatePresence>
        {tiles.map((tile, index) => {
          return (
            <Tile key={index} className="grid-cell" value={tile}/>
          );
        })}
      </AnimatePresence>
    </div>
  );
};
