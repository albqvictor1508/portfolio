import { createContext, type ReactNode, useContext, useState } from "react";
import en from "../translations/en.json";
import pt from "../translations/pt.json";

type Language = "en" | "pt";

type Translations = typeof en;

interface LanguageContextType {
  language: Language;
  setLanguage: (lang: Language) => void;
  toggleLanguage: () => void;
  t: (key: string) => string;
}

const LanguageContext = createContext<LanguageContextType | undefined>(
  undefined,
);

const translations: Record<Language, Translations> = {
  en: en as Translations,
  pt: pt as Translations,
};

export const LanguageProvider = ({ children }: { children: ReactNode }) => {
  const [language, setLanguage] = useState<Language>("pt");

  const t = (key: string): string => {
    const currentTranslations = translations[language];
    const keys = key.split(".");
    let result: any = currentTranslations;
    for (const k of keys) {
      if (!result || typeof result !== "object" || !(k in result)) {
        return key;
      }
      result = result[k];
    }
    return typeof result === "string" ? result : key;
  };

  const toggleLanguage = () => {
    setLanguage((prevLang) => (prevLang === "en" ? "pt" : "en"));
  };

  return (
    <LanguageContext.Provider value={{ language, setLanguage, toggleLanguage, t }}>
      {children}
    </LanguageContext.Provider>
  );
};

export const useLanguage = () => {
  const context = useContext(LanguageContext);
  if (context === undefined) {
    throw new Error("useLanguage must be used within a LanguageProvider");
  }
  return context;
};
