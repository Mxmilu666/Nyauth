import { useMessageDialog } from '@/components/service/index'

export const message = {
    info(message: string) {
        return useMessageDialog(message)
    },

    success(message: string) {
        return useMessageDialog(message, 'success')
    },

    error(message: string) {
        return useMessageDialog(message, 'error')
    },

    warning(message: string) {
        return useMessageDialog(message, 'warning')
    }
}
