import { inject, computed } from 'vue'

export function useConsole() {
	const avatar = inject('avatar') as string
	const currentHour = new Date().getHours()
	const timeGreeting = computed(() => {
		if (currentHour < 5) return '时间不早了哦，该休息啦'
		if (currentHour < 9) return '早哦~ 要不要再继续睡一会zzz'
		if (currentHour < 12) return '中午好~ 该吃午饭啦'
		if (currentHour < 18) return '下午好哦，要不要来点下午茶？'
		return '晚上好~ 吃过晚饭了嘛?'
	})
	return { avatar, timeGreeting }
}
