import { useMountComponent } from '@/hooks/useMountComponent'

import MessageComponent from './Message.vue'

export const useMessageDialog = async (message: string) => {
    return await useMountComponent({ message }).mount(MessageComponent)
}
