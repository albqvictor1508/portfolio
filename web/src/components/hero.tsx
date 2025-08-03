export const HeroSection = () => {
	return (
		<div className="w-full h-full flex flex-col gap-4">
			<div className="w-full p-4 flex gap-4 border border-blue-900  rounded-md">
				<div className="flex flex-col gap-4 flex-1 justify-center items-center">
					<div className="w-full">
						<h1>hey, I'm Victor</h1>
						<p>Lorem</p>
					</div>
					<div className="flex justify-center items-center w-full h-[200px] rounded-xl border border-zinc-700">
						salve
					</div>
				</div>

				<div className="flex flex-1 justify-center items-center">
					<h1>hey, I'm Victor</h1>
					<p>Lorem</p>
				</div>
			</div>

			<div className="w-full flex justify-between">
				<span>salve</span>
				<span>salve</span>
				<span>salve</span>
				<span>salve</span>
			</div>
		</div>
	);
};
