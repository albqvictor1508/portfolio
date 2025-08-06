import { Badge } from "./ui/badge";

export const StackSection = () => {
  const technologies = [
    { icon: "/icons/go.svg", name: "Go" },
    { icon: "/icons/java.svg", name: "Gin" },
    { icon: "/icons/typescript.svg", name: "Node.js" },
    { icon: "/icons/mongodb.svg", name: "Bun" },
    { icon: "/icons/mongodb.svg", name: "Git" },
    { icon: "/icons/java.svg", name: "Java" },
    { icon: "/icons/java.svg", name: "Springboot" },
    { icon: "/icons/docker.svg", name: "Docker" },
    { icon: "/icons/aws.svg", name: "AWS" },
    { icon: "/icons/react.svg", name: "React Native" },
    { icon: "/icons/postgres.svg", name: "Postgres" },
    { icon: "/icons/mongodb.svg", name: "MongoDB" },
    { icon: "/icons/mongodb.svg", name: "Express.js" },
    { icon: "/icons/mongodb.svg", name: "Fastify" },
    { icon: "/icons/mongodb.svg", name: "Elysia" },
    { icon: "/icons/websocket.svg", name: "Websocket" },
  ];
  return (
    <div className="w-full h-full flex flex-col gap-4">
      <h2 className="text-2xl font-semibold">My Technologies.</h2>
      <p className="text-sm text-zinc-400">
        I work with a modern stack, focusing on performance, scalability, and
        robust solutions.
      </p>
      <div className="flex m-auto gap-3 items-center flex-wrap ">
        {technologies.map((tech: { icon: string; name: string }) => (
          <Badge
            variant={"outline"}
            className="p-2 flex justify-center items-center gap-2"
            key={tech.name}
          >
            <span className="w-6 h-6">
              <img
                src={tech.icon}
                className="w-full h-full"
                alt={`${tech.name} icon`}
              />
            </span>
            {tech.name}
          </Badge>
        ))}
      </div>
    </div>
  );
};

