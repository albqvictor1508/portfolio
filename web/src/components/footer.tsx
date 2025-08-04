export const Footer = ({ isLight }: { isLight: boolean }) => {
	const textStyle = isLight ? "text-zinc-400" : "text-zinc-800";
	return (
		<div className="w-full flex justify-center items-center">
			<div className="flex justify-center items-center border-2 border-orange-400 w-1/2 p-4 ">
				<div className="flex flex-1 flex-col gap-2">
					<h3 className="font-semibold text-lg">Other</h3>
					<span className={`${textStyle} hover:underline cursor-pointer`}>
						Home
					</span>
					<span className={`${textStyle} hover:underline cursor-pointer`}>
						Home
					</span>
					<span className={`${textStyle} hover:underline cursor-pointer`}>
						Home
					</span>
				</div>
				<div className="flex flex-1 flex-col gap-2">
					<h3 className="font-semibold text-lg">Other</h3>
					<span className={`${textStyle} hover:underline cursor-pointer`}>
						Home
					</span>
					<span className={`${textStyle} hover:underline cursor-pointer`}>
						Home
					</span>
					<span className={`${textStyle} hover:underline cursor-pointer`}>
						Home
					</span>
				</div>
				<div className="flex flex-1 flex-col gap-2">
					<h3 className="font-semibold text-lg">Other</h3>
					<span className={`${textStyle} hover:underline cursor-pointer`}>
						Home
					</span>
					<span className={`${textStyle} hover:underline cursor-pointer`}>
						Home
					</span>
					<span className={`${textStyle} hover:underline cursor-pointer`}>
						Home
					</span>
				</div>
			</div>
		</div>
	);
};
