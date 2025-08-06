export const Footer = ({ isLight }: { isLight: boolean }) => {
  const textStyle = isLight ? "text-zinc-400" : "text-zinc-800";
  return (
    <div
      id="footer"
      className="w-2/3 border-2 border-rose-500 flex justify-center items-center"
    >
      <div className="w-full flex p-4 border-2 border-blue-500 justify-center">
        <div className="flex flex-1 flex-col justify-center gap-2">
          <h3 className="w-fit font-semibold text-base border border-green-400">
            Important Links
          </h3>
          <span
            className={`${textStyle} text-sm w-fit hover:underline cursor-pointer`}
          >
            Home
          </span>
          <span
            className={`${textStyle} text-sm w-max border border-red-500 hover:underline cursor-pointer`}
          >
            Home
          </span>
          <span
            className={`${textStyle} text-sm w-max hover:underline cursor-pointer`}
          >
            Home
          </span>
        </div>
        <div className="flex flex-1 flex-col gap-2 justify-center">
          <h3 className="w-max font-semibold text-base">Other</h3>
          <span
            className={`${textStyle} text-sm w-max hover:underline cursor-pointer`}
          >
            Home
          </span>
          <span
            className={`${textStyle} text-sm w-max hover:underline cursor-pointer`}
          >
            Home
          </span>
          <span
            className={`${textStyle} text-sm w-max hover:underline cursor-pointer`}
          >
            Home
          </span>
        </div>
        <div className="flex flex-1 flex-col gap-2 justify-center">
          <h3 className="w-max font-semibold text-base">Other</h3>
          <span
            className={`${textStyle} text-sm w-max hover:underline cursor-pointer`}
          >
            Home
          </span>
          <span
            className={`${textStyle} text-sm w-max hover:underline cursor-pointer`}
          >
            Home
          </span>
          <span
            className={`${textStyle} text-sm w-max hover:underline cursor-pointer`}
          >
            Home
          </span>
        </div>
      </div>
    </div>
  );
};
