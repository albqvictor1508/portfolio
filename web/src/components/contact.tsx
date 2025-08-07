import { useState } from "react";
import { z } from "zod";
import { Button } from "./ui/button";
import { Input } from "./ui/input";
import { Textarea } from "./ui/textarea.tsx";

const contactFormSchema = z.object({
  name: z.string().min(1, "Name is required."),
  email: z.string().email("Invalid email address."),
  message: z.string().min(1, "Message is required.").max(500),
});

type ContactFormData = z.infer<typeof contactFormSchema>;

export const ContactSection = () => {
  const [formData, setFormData] = useState<Partial<ContactFormData>>({
    name: "",
    email: "",
    message: "",
  });
  const [errors, setErrors] = useState<
    Partial<Record<keyof ContactFormData, string>>
  >({});

  const handleChange = (
    e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>,
  ) => {
    const { name, value } = e.target;
    setFormData((prev) => ({ ...prev, [name]: value }));
  };

  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    const result = contactFormSchema.safeParse(formData);

    if (!result.success) {
      const newErrors: Partial<Record<keyof ContactFormData, string>> = {};
      for (const issue of result.error.issues) {
        newErrors[issue.path[0] as keyof ContactFormData] = issue.message;
      }
      setErrors(newErrors);
      return;
    }

    setErrors({});
    alert("Thank you for your message!");
    setFormData({ name: "", email: "", message: "" });
  };

  const messageLength = formData.message?.length || 0;

  return (
    <div className="w-full h-full flex flex-col gap-4">
      <h2 className="text-lg font-semibold">Contact me.</h2>
      <p className="text-sm text-zinc-400">lorem ipsum.</p>

      <form
        onSubmit={handleSubmit}
        className="w-full h-full flex flex-col p-4 gap-8 rounded-md bg-zinc-900"
      >
        <div
          className="w-full flex gap-6
          "
        >
          <div className="w-full">
            <Input
              name="name"
              value={formData.name}
              onChange={handleChange}
              placeholder="Your name."
              aria-invalid={!!errors.name}
            />
            {errors.name && (
              <p className="text-xs text-red-500 mt-1">{errors.name}</p>
            )}
          </div>
          <div className="w-full">
            <Input
              name="email"
              type="email"
              value={formData.email}
              onChange={handleChange}
              placeholder="Your email."
              aria-invalid={!!errors.email}
            />
            {errors.email && (
              <p className="text-xs text-red-500 mt-1">{errors.email}</p>
            )}
          </div>
        </div>
        <div className="w-full flex flex-col gap-2">
          <Textarea
            name="message"
            value={formData.message}
            onChange={handleChange}
            className="w-full h-[150px]"
            maxLength={500}
            minLength={3}
            placeholder="Your message"
            aria-invalid={!!errors.message}
          />
          {errors.message && (
            <p className="text-xs text-red-500">{errors.message}</p>
          )}
          <div className="text-xs text-zinc-400 w-full flex justify-end">
            <span>{messageLength}/500 characters</span>
          </div>
        </div>
        <div className="w-full flex justify-end">
          <Button type="submit">Send Message</Button>
        </div>
      </form>
    </div>
  );
};
