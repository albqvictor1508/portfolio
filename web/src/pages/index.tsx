import { createFileRoute } from "@tanstack/react-router";
import { AboutSection } from "@/components/about";
import { ContactSection } from "@/components/contact";
import { ExperienceSection } from "@/components/experience";
import { Footer } from "@/components/footer";
import { HeroSection } from "@/components/hero";
import { Menu } from "@/components/menu";
import { ProjectSection } from "@/components/project/section";
import { StackSection } from "@/components/stack";

export const Route = createFileRoute("/")({
  component: RouteComponent,
});

function RouteComponent() {
  return (
    <div className="w-[900px] h-full flex flex-col gap-20 m-auto justify-center items-center">
      <Menu />
      <HeroSection />
      <AboutSection />
      <ProjectSection />
      <ExperienceSection />
      <StackSection />
      <ContactSection />
      <Footer />
    </div>
  );
}
