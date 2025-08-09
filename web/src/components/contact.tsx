import { Loader2 } from "lucide-react";
import { useState } from "react";
import { z } from "zod";
import { useLanguage } from "../context/LanguageContext";
import { Button } from "./ui/button";
import { Input } from "./ui/input";
import { Textarea } from "./ui/textarea.tsx";

const contactFormSchema = z.object({
  name: z.string().min(1, "Name is required."),
  email: z.email("Invalid email address."),
  message: z.string().min(1, "Message is required.").max(500),
});

type ContactFormData = z.infer<typeof contactFormSchema>;

export const ContactSection = () => {
  const { t } = useLanguage();
  const [formData, setFormData] = useState<Partial<ContactFormData>>({
    name: "",
    email: "",
    message: "",
  });
  const [errors, setErrors] = useState<
    Partial<Record<keyof ContactFormData, string>>
  >({});

  const [isLoading, setIsLoading] = useState(false);

  const [submitStatus, setSubmitStatus] = useState<"success" | "error" | null>(
    null,
  );

  const handleChange = (
    e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>,
  ) => {
    const { name, value } = e.target;
    setFormData((prev) => ({ ...prev, [name]: value }));
  };

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
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
    setIsLoading(true);
    setSubmitStatus(null);
    try {
      const response = await fetch("http://localhost:3333/contact", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          user_email: formData.email,
          subject: `${formData.name} - By Portfolio`,
          content: formData.message,
        }),
      });

      if (response.ok) {
        setSubmitStatus("success");
        setFormData({ name: "", email: "", message: "" });
      } else {
        setSubmitStatus("error");
      }
    } catch (error) {
      console.error("Error sending message:", error);
      setSubmitStatus("error");
    } finally {
      setIsLoading(false);
    }
  };

  const messageLength = formData.message?.length || 0;

  return (
    <div className="w-full h-full flex flex-col gap-4">
      <h2 className="text-2xl font-semibold">{t("contact_section.title")}</h2>
      <p className="text-sm text-zinc-400 leading-6">
        {t("contact_section.description")}
      </p>

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
              placeholder={t("contact_section.your_name")}
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
              placeholder={t("contact_section.your_email")}
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
            className="w-full h-[150px] resize-none"
            maxLength={500}
            minLength={3}
            placeholder={t("contact_section.your_message")}
            aria-invalid={!!errors.message}
          />
          {errors.message && (
            <p className="text-xs text-red-500">{errors.message}</p>
          )}
          <div className="text-xs text-zinc-400 w-full flex justify-end">
            <span>
              {messageLength}/500 {t("contact_section.characters")}
            </span>
          </div>
        </div>
        <div className="w-full flex justify-end items-center gap-4">
          {submitStatus === "success" && (
            <p className="text-xs text-green-500">
              {t("contact_section.success_message")}
            </p>
          )}
          {submitStatus === "error" && (
            <p className="text-xs text-red-500">Failed to send message.</p>
          )}
          <Button
            type="submit"
            disabled={isLoading}
            className="w-40 h-10 flex items-center justify-center"
          >
            {isLoading ? (
              <Loader2 className="animate-spin" />
            ) : (
              t("contact_section.send_message")
            )}
          </Button>
        </div>
      </form>
    </div>
  );
};
