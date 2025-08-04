export const Footer = ({ isLight }: { isLight: boolean }) => {
  const textStyle = isLight ? "text-zinc-400" : "text-zinc-800";
  return (
    <div className="w-full flex justify-center items-center">
      <div className="w-3/4 flex p-4 ">
        <div className="flex flex-1 flex-col gap-2">
          <h3 className="text-center w-max font-semibold text-lg">Other</h3>
          <span
            className={`${textStyle} w-max text-center hover:underline cursor-pointer`}
          >
            Home
          </span>
          <span
            className={`${textStyle} w-max text-center hover:underline cursor-pointer`}
          >
            Home
          </span>
          <span
            className={`${textStyle} w-max text-center hover:underline cursor-pointer`}
          >
            Home
          </span>
        </div>
        <div className="flex flex-1 flex-col gap-2">
          <h3 className="text-center w-max font-semibold text-lg">Other</h3>
          <span
            className={`${textStyle} w-max text-center hover:underline cursor-pointer`}
          >
            Home
          </span>
          <span
            className={`${textStyle} w-max text-center hover:underline cursor-pointer`}
          >
            Home
          </span>
          <span
            className={`${textStyle} w-max text-center hover:underline cursor-pointer`}
          >
            Home
          </span>
        </div>
        <div className="flex flex-1 flex-col gap-2">
          <h3 className="text-center w-max font-semibold text-lg">Other</h3>
          <span
            className={`${textStyle} w-max text-center hover:underline cursor-pointer`}
          >
            Home
          </span>
          <span
            className={`${textStyle} w-max text-center hover:underline cursor-pointer`}
          >
            Home
          </span>
          <span
            className={`${textStyle} w-max text-center hover:underline cursor-pointer`}
          >
            Home
          </span>
        </div>
      </div>
    </div>
  );
};
