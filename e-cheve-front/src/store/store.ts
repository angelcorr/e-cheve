import { ref } from 'vue'

export const store = ref({
  emails: [
    {
      id: 12,
      from: 'cheme@gmail.com',
      to: 'chumi@gmail.com',
      date: new Date().toLocaleDateString(),
      subject: 'Remember me',
      body: 'Remember me, please. Remember you, please, And my treats?. I have a orange mark:). I have a orange mark:) Remember me, please. Remember you, please, And my treats?. I have a orange mark:). I have a orange mark:) Remember me, please. Remember you, please, And my treats?. I have a orange mark:). I have a orange mark:) Remember me, please. Remember you, please, And my treats?. I have a orange mark:). I have aorange mark:)Remember me, please. Remember you, please, And my treats?. I have a orange mark:). I have aorange mark:)Remember me, please. Remember you, please, And my treats?. I have a orange mark:). I have aorange mark:)Remember me, please. Remember you, please, And my treats?. I have a orange mark:). I have aorange mark:)Remember me, please. Remember you, please, And my treats?. I have a orange mark:). I have aorange mark:)Remember me, please. Remember you, please, And my treats?. I have a orange mark:). I have aorange mark:)Remember me, please. Remember you, please, And my treats?. I have a orange mark:). I have aorange mark:)Remember me, please. Remember you, please, And my treats?. I have a orange mark:). I have aorange mark:)Remember me, please. Remember you, please, And my treats?. I have a orange mark:). I have aorange mark:)Remember me, please. Remember you, please, And my treats?. I have a orange mark:). I have aorange mark:)'
    },
    {
      id: 23,
      from: 'chumi@gmail.com',
      to: 'cheme@gmail.com',
      date: new Date().toLocaleDateString(),
      subject: 'Remember you',
      body: 'Remember you, please'
    },
    {
      id: 34,
      from: 'yoyi@gmail.com',
      to: 'yonyin@gmail.com',
      date: new Date().toLocaleDateString(),
      subject: 'I want treats',
      body: 'And my treats?'
    },
    {
      id: 45,
      from: 'drondrin@gmail.com',
      to: 'gordon@gmail.com',
      date: new Date().toLocaleDateString(),
      subject: 'I am orange',
      body: 'I have a orange mark:)'
    },
    {
      id: 56,
      from: 'drondrin@gmail.com',
      to: 'gordon@gmail.com',
      date: new Date().toLocaleDateString(),
      subject: 'I am orange',
      body: 'I have a orange mark:)'
    }
  ]
})
