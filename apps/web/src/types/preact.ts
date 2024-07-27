import type { ComponentChildren } from "preact";

export type PropsWithChildren<T = {}> = T & {
  children?: ComponentChildren;
};
