import { useForm } from "react-hook-form";
import { Input } from "./ui/input";
import { Textarea } from "./ui/textarea.tsx";

export const ContactSection = () => {
  const { register, watch } = useForm<{
    name: string;
    email: string;
    message: string;
  }>();

  const message = watch("message", "");

  return (
    <div className="w-full h-full flex flex-col gap-4">
      <h2 className="text-lg font-semibold">Contact me.</h2>
      <p className="text-sm text-zinc-400">lorem ipsum.</p>

      <form className="w-full h-full flex flex-col p-4 gap-8 rounded-md bg-zinc-900">
        <div
          className="w-full flex gap-6
          "
        >
          <Input {...register("name")} placeholder="Your name." />
          <Input {...register("email")} placeholder="Your email." />
        </div>
        <div className="w-full flex flex-col gap-2">
          <Textarea
            {...register("message")}
            className="w-full h-[150px]"
            maxLength={500}
            minLength={3}
            placeholder="Your message"
          />
          <div className="text-xs text-zinc-400 w-full flex justify-end">
            <span>{message.length}/500 characters</span>
          </div>
        </div>
      </form>
    </div>
  );
};
