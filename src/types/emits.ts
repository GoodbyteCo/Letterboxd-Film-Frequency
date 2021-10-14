export type ControlEmits = {
  (event: "changeUsername", value: string): void;
  (event: "changeYear", value: number): void;
};
