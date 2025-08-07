import { createFileRoute } from "@tanstack/react-router";
import { AboutSection } from "@/components/about";
import { ContactSection } from "@/components/contact";
import { ExperienceSection } from "@/components/experience";
import { Footer } from "@/components/footer";
import { HeroSection } from "@/components/hero";
import { Menu } from "@/components/menu";
import { ProjectSection } from "@/components/project/section";
import { StackSection } from "@/components/stack";
import { useEffect, useState } from "react";

export const Route = createFileRoute("/")({
  component: RouteComponent,
});

function RouteComponent() {
  const [commits, setCommits] = useState([]);

  useEffect((): void => {
    const fetchData = async () => {
      const reply = await fetch(
        "http://localhost:3333/commits?since=2025-06-01",
        {
          method: "GET",
        },
      );
      const data = await reply.json();
      console.log(data);
      setCommits(data);
    };
    fetchData();
  }, []);

  return (
    <div className="w-[1100px] h-full flex flex-col gap-14 m-auto justify-center items-center">
      <Menu />
      <HeroSection />
      <AboutSection commits={commits} />
      <ProjectSection />
      <ExperienceSection />
      <StackSection />
      <ContactSection />
      <Footer isLight={true} />
    </div>
  );
}
