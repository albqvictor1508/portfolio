export const Footer = ({ isLight }: { isLight: boolean }) => {
  const textStyle = isLight ? "text-zinc-400" : "text-zinc-800";
  return (
    <div id="footer" className="w-2/3 flex justify-center items-center">
      <div className="w-full flex p-4 justify-around">
        <div className="flex flex-col justify-center gap-2">
          <h3 className="w-fit font-semibold text-base">Important Links</h3>
          <span
            className={`${textStyle} text-sm w-fit hover:underline cursor-pointer`}
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
        <div className="flex  flex-col gap-2 justify-center">
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
        <div className="w-max flex  flex-col gap-2 justify-center">
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
