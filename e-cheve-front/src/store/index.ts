import { reactive } from 'vue'

export interface Email {
  id: string
  from: string
  to: string
  date: string
  subject: string
  body: string
}

export const store = reactive({
  emails: [] as Email[],
  term: '',
  emailSelected: {},
  changeEmailSelected(email: Email) {
    this.emailSelected = email;
    return this.emailSelected;
  }
})
