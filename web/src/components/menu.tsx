import { Link } from "@tanstack/react-router";
import { CircleCheckIcon, CircleHelpIcon, CircleIcon } from "lucide-react";

import {
  NavigationMenu,
  NavigationMenuContent,
  NavigationMenuItem,
  NavigationMenuLink,
  NavigationMenuList,
  NavigationMenuTrigger,
  navigationMenuTriggerStyle,
} from "@/components/ui/navigation-menu";
import { useLanguage } from "../context/LanguageContext";

export function Menu() {
  const { setLanguage, t } = useLanguage();

  return (
    <div className="w-full mt-10 flex justify-between items-center">
      <div className="flex items-center gap-4">
        <Link to=".." className="font-bold">
          VA
        </Link>
        <NavigationMenu viewport={false}>
          <NavigationMenuList>
            <NavigationMenuItem>
              <NavigationMenuLink
                asChild
                className={navigationMenuTriggerStyle()}
              >
                <a href="/">{t("menu.home")}</a>
              </NavigationMenuLink>
            </NavigationMenuItem>

            <NavigationMenuItem>
              <NavigationMenuLink
                asChild
                className={navigationMenuTriggerStyle()}
              >
                <a href="/about">{t("menu.about")}</a>
              </NavigationMenuLink>
            </NavigationMenuItem>

            <NavigationMenuItem>
              <NavigationMenuTrigger>
                {t("menu.with_icon")}
              </NavigationMenuTrigger>
              <NavigationMenuContent>
                <ul className="grid w-[200px] gap-4">
                  <li>
                    <NavigationMenuLink asChild>
                      <a href="#" className="flex-row items-center gap-2">
                        <CircleHelpIcon />
                        {t("menu.github")}
                      </a>
                    </NavigationMenuLink>
                    <NavigationMenuLink asChild>
                      <a href="#" className="flex-row items-center gap-2">
                        <CircleIcon />
                        {t("menu.linkedin")}
                      </a>
                    </NavigationMenuLink>
                    <NavigationMenuLink asChild>
                      <a href="#" className="flex-row items-center gap-2">
                        <CircleCheckIcon />
                        {t("menu.done")}
                      </a>
                    </NavigationMenuLink>
                  </li>
                </ul>
              </NavigationMenuContent>
            </NavigationMenuItem>
          </NavigationMenuList>
        </NavigationMenu>
      </div>
      <NavigationMenu viewport={false}>
        <NavigationMenuList>
          <NavigationMenuItem>
            <NavigationMenuTrigger className="text-white">
              🌐
            </NavigationMenuTrigger>
            <NavigationMenuContent>
              <ul className="grid w-[100px] gap-1 p-2">
                <li>
                  <NavigationMenuLink asChild>
                    <a
                      href="#"
                      className="flex items-center gap-2"
                      onClick={() => setLanguage("pt")}
                    >
                      🇧🇷 PT
                    </a>
                  </NavigationMenuLink>
                </li>
                <li>
                  <NavigationMenuLink asChild>
                    <a
                      href="#"
                      className="flex items-center gap-2"
                      onClick={() => setLanguage("en")}
                    >
                      🇺🇸 EN
                    </a>
                  </NavigationMenuLink>
                </li>
              </ul>
            </NavigationMenuContent>
          </NavigationMenuItem>
        </NavigationMenuList>
      </NavigationMenu>
    </div>
  );
}
