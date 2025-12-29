import { useDrag } from "@use-gesture/react";
import { useEffect } from "react";
import { Board } from "./components/Board";
import { useWasmGame } from "./hooks/useWasmGame";
import "./App.css";

function App() {
	const { board, move, newGame, isReady } = useWasmGame();

	// Keyboard controls
	useEffect(() => {
		const handleKeyDown = (e) => {
			switch (e.key) {
				case "ArrowUp":
					move("up");
					break;
				case "ArrowDown":
					move("down");
					break;
				case "ArrowLeft":
					move("left");
					break;
				case "ArrowRight":
					move("right");
					break;
			}
		};
		window.addEventListener("keydown", handleKeyDown);
		return () => window.removeEventListener("keydown", handleKeyDown);
	}, [move]);

	// Gestures
	const bind = useDrag(({ swipe: [swipeX, swipeY] }) => {
		// swipe is [0,0] if not swipe, or [-1 | 1, -1 | 1]
		if (swipeX === -1) move("left");
		else if (swipeX === 1) move("right");
		else if (swipeY === -1) move("up");
		else if (swipeY === 1) move("down");
	});

	return (
		<div
			className="app-container"
			{...bind()}
			style={{
				touchAction: "none",
				height: "100vh",
				display: "flex",
				flexDirection: "column",
				alignItems: "center",
				justifyContent: "center",
			}}
		>
			<div
				className="glass-panel"
				style={{
					padding: "2rem",
					display: "flex",
					flexDirection: "column",
					alignItems: "center",
					gap: "1.5rem",
				}}
			>
				<h1 className="title">2048</h1>

				<div
					className="score-board glass-panel"
					style={{ padding: "10px 20px", borderRadius: "12px" }}
				>
					<span
						style={{
							fontSize: "0.8rem",
							textTransform: "uppercase",
							letterSpacing: "1px",
						}}
					>
						Score
					</span>
					<div style={{ fontSize: "1.5rem", fontWeight: "bold" }}>
						{board.Score}
					</div>
				</div>

				{!isReady && <div>Loading Game...</div>}

				{isReady && <Board tiles={board.Tiles || []} />}

				<div className="instructions">Use Arrow Keys or Swipe to Play</div>

				<button
					type="button"
					onClick={newGame}
					className="glass-panel new-game-button"
				>
					New Game
				</button>
			</div>
		</div>
	);
}

export default App;
