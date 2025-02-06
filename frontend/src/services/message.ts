import { useMessageDialog } from '@/components/service/index'

export const message = {
    info(message: string) {
        return useMessageDialog(message)
    }
}
