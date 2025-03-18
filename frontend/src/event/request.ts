const eventList = ['Unauthorized', 'NetworkError', 'UnknownError', 'Message'] as const

export type EventList = (typeof eventList)[number]

const _listenClient = new Map<string, Set<Function>>()

export const on = (event: EventList, callback: Function) => {
    if (!_listenClient.has(event)) _listenClient.set(event, new Set<Function>())
    _listenClient.get(event)?.add(callback)
}

export const emit = (event: EventList, ...args: any[]) => {
    _listenClient.get(event)?.forEach((callback) => {
        callback(...args)
    })
}

const requestEvent = {
    on,
    emit,
    eventList
}

export default requestEvent
