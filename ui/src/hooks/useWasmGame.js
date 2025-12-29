import { useEffect, useRef, useState } from "react";

export function useWasmGame() {
	const [board, setBoard] = useState({ Score: 0, Tiles: [], GameOver: false });
	const [isReady, setIsReady] = useState(false);
	const goRef = useRef(null);

	useEffect(() => {
		const loadWasm = async () => {
			if (!window.Go) {
				console.error("wasm_exec.js not loaded");
				return;
			}
			const go = new window.Go();
			goRef.current = go;

			try {
				const result = await WebAssembly.instantiateStreaming(
					fetch("/main.wasm"),
					go.importObject,
				);
				go.run(result.instance);
				setIsReady(true);

				// Initialize game
				if (window.newGame) {
					const stateStr = window.newGame();
					setBoard(JSON.parse(stateStr));
				}
			} catch (e) {
				console.error("Failed to load WASM", e);
			}
		};

		loadWasm();
	}, []);

	const move = (direction) => {
		if (!isReady || !window.move) return;
		const stateStr = window.move(direction);
		setBoard(JSON.parse(stateStr));
	};

	const newGame = () => {
		if (!isReady || !window.newGame) return;
		const stateStr = window.newGame();
		setBoard(JSON.parse(stateStr));
	};

	return { board, move, newGame, isReady };
}
