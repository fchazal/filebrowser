<template>
  <article id="settings">
  <form @submit.prevent="save">
    <header>
      <h1>{{ $t('settings.profileSettings') }}</h1>

      <button type="submit" class="action">
        <i class="material-icons">save</i>
      </button>
    </header>
  
    <main>
      <fieldset>
        <user-form :user.sync="me" :isDefault="false" :isNew="false" />
      </fieldset>
    </main>

<!--
    <form @submit="updateSettings">
      <h2>{{ $t('settings.profileSettings') }}</h2>
      <button type="submit" class="action">
        <i class="material-icons">check</i>
      </button>

      <div class="content">
        <p><input type="checkbox" v-model="hideDotfiles"> {{ $t('settings.hideDotfiles') }}</p>
        <p><input type="checkbox" v-model="singleClick"> {{ $t('settings.singleClick') }}</p>
        
        <h3>{{ $t('settings.language') }}</h3>
        <languages class="input input--block" :locale.sync="locale"></languages>
      </div>
    </form>

    <form v-if="!user.lockPassword" @submit="updatePassword">
      <h2>{{ $t('settings.changePassword') }}</h2>
      <button type="submit" class="action">
        <i class="material-icons">check</i>
      </button>

      <div class="content">
        <input :class="passwordClass" type="password" :placeholder="$t('settings.newPassword')" v-model="password" name="password">
        <input :class="passwordClass" type="password" :placeholder="$t('settings.newPasswordConfirm')" v-model="passwordConf" name="password">
      </div>
    </form>
-->
  </form>
  </article>
</template>

<script>
import { mapState, mapMutations } from 'vuex'
import { users as api } from '@/api'
import UserForm from '@/components/settings/UserForm'

export default {
  name: 'settings',
  components: {
    UserForm
  },
  data: function () {
    return {
      password: '',
      passwordConf: '',
      hideDotfiles: false,
      singleClick: false,
      locale: '',
      me: {}
    }
  },
  computed: {
    ...mapState([ 'user' ]),
    passwordClass () {
      const baseClass = 'input input--block'

      if (this.password === '' && this.passwordConf === '') {
        return baseClass
      }

      if (this.password === this.passwordConf) {
        return `${baseClass} input--green`
      }

      return `${baseClass} input--red`
    }
  },
  created () {
    this.locale = this.user.locale
    this.hideDotfiles = this.user.hideDotfiles
    this.singleClick = this.user.singleClick
    this.fetchData()
  },
  methods: {
    ...mapMutations([ 'updateUser' ]),
     async fetchData () {
      const id = this.user.id
      this.me = { ...await api.get(id) }
      console.log(this.me)
    },
    async updatePassword (event) {
      event.preventDefault()

      if (this.password !== this.passwordConf || this.password === '') {
        return
      }

      try {
        const data = { id: this.user.id, password: this.password }
        await api.update(data, ['password'])
        this.updateUser(data)
        this.$showSuccess(this.$t('settings.passwordUpdated'))
      } catch (e) {
        this.$showError(e)
      }
    },
    async updateSettings (event) {
      event.preventDefault()

      try {
        const data = { id: this.user.id, locale: this.locale, hideDotfiles: this.hideDotfiles, singleClick: this.singleClick }
        await api.update(data, ['locale', 'hideDotfiles', 'singleClick'])
        this.updateUser(data)
        this.$showSuccess(this.$t('settings.settingsUpdated'))
      } catch (e) {
        this.$showError(e)
      }
    }
  }
}
</script>
