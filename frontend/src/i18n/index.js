import Vue from 'vue'
import VueI18n from 'vue-i18n'

import en from './en.json'
import fr from './fr.json'

Vue.use(VueI18n)

export function detectLocale () {
  let locale = (navigator.language || navigator.browserLangugae).toLowerCase()
  switch (true) {
    case /^fr.*/i.test(locale):
      locale = 'fr'
      break
    default:
      locale = 'en'
  }

  return locale
}

const i18n = new VueI18n({
  locale: detectLocale(),
  fallbackLocale: 'en',
  messages: {
    'en': en,
    'fr': fr,
  }
})

export default i18n
