import TypeWriter from "typewriter-effect";

type TypingEffectProps = {
  text: string;
};

export const TypingEffect = ({ text }: TypingEffectProps) => {
  return (
    <div>
      <TypeWriter
        options={{
          strings: [text, text, text, text, text, text],
          autoStart: true,
          loop: true,
        }}
      />
    </div>
  );
};
