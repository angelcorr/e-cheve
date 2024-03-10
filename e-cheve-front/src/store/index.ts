import { ref } from 'vue'

export interface Email {
  id: string
  from: string
  to: string
  date: string
  subject: string
  body: string
}

export const store = ref({
  emails: [] as Email[],
  term: ''
})
