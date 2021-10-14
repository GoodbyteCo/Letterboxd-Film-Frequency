export type ControlsProps = {
    lowestYear?: number
}

export type GraphProps = {
    year: number,
    films: Record<string, Record<string, number>>,
    username: string
}

export type StatusProps = { type: string; message: string; }