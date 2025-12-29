import React from 'react';
import {motion} from 'framer-motion';

export const Tile = ({className, value}) => {
  const displayVal = Math.pow(2, value);

  return (
    <motion.div
      key={value}
      initial={{scale: 0, opacity: 0}}
      animate={{scale: 1, opacity: 1}}
      exit={{scale: 0, opacity: 0}}
      transition={{type: "spring", stiffness: 400, damping: 25}}
      className={`tile ${className}`}
      style={{
        backgroundColor: `var(--tile-${displayVal}, rgba(255, 255, 255, 0.4))`,
      }}
    >
      <div className="tile-inner">
        {value > 0 ? displayVal : null}
      </div>
    </motion.div>
  );
};
