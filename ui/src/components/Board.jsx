import { AnimatePresence } from "framer-motion";
import { Tile } from "./Tile";

export const Board = ({ tiles }) => (
	<div className="game-grid">
		<AnimatePresence>
			{tiles.map((tile, index) => (
				// biome-ignore lint/suspicious/noArrayIndexKey: <explanation>
				<Tile key={index} className="grid-cell" value={tile} />
			))}
		</AnimatePresence>
	</div>
);
