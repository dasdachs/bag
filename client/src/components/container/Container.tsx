import React, { ReactChild } from "react";
import { useStyletron } from "baseui";

export function Container({ children }: { children: ReactChild }) {
  const [css] = useStyletron();
  return (
    <section
      className={css({
        alignContent: "center",
        display: "flex",
        justifyContent: "center",
        marginBottom: "16px",
        marginTop: "16px",
      })}
    >
      <div className={css({ width: "90%" })}>{children}</div>
    </section>
  );
}
