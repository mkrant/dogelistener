export type DogeSession = {
    id: string,
    is_running: boolean,
}

export type SessionDetail = {
    id: string,
    is_running: boolean,
    runs: Run[],
}

export type Run = {
    id: string,
    start_time: string,
    duration_seconds: number,
    live: boolean,
}

export type RunDetail = {
    id: string,
    start_time: string,
    duration_seconds: number,
    live: boolean,
    data: number[],
}