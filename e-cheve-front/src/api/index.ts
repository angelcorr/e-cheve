import { store, type Email } from '../store'

interface ApiHit {
  _id: string
  _source: {
    From: string
    To: string
    Date: string
    Subject: string
    Body: string
  }
}

interface ApiResponse {
  hits: {
    hits: ApiHit[]
  }
}

export const getEmails = async (pageSize: boolean) => {
  store.loader(true);
  if (store.term.length < 1) {
    store.loader(false);
    return;
  }

  let page;

  if (!pageSize) {
    page = 0;
  } else {
    page = store.emails.length - 1;
  }

  const endpoint = `${import.meta.env.VITE_API_BASE_URL}/emails?term=${store.term}&page=${page}`
  const response = await fetch(endpoint)
  if (response.status !== 200) {
    store.loader(false);
    return 'There was a problem' //To improve
  } 

  const apiResponse: ApiResponse = await response.json()

  const emails: Email[] = apiResponse.hits.hits.map((hit) => ({
    id: hit._id,
    from: hit._source.From,
    to: hit._source.To,
    date: hit._source.Date,
    subject: hit._source.Subject,
    body: hit._source.Body
  }))
  store.loader(false);

  if (!pageSize) {
    store.emails = [...emails]
  } else {
    store.emails = [...store.emails, ...emails]
  }
}
