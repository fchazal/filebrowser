<template>
  <article id="settings" v-if="settings !== null">
  <form v-if="loaded" @submit.prevent="save">
    <header>
      <button @click.prevent="$router.back()" class="action">
        <i class="material-icons">arrow_back</i>
      </button>

      <h1 v-if="user.id === 0">{{ $t('settings.newUser') }}</h1>
      <h1 v-else>{{ $t('settings.user') }} <i>{{ user.username }}</i></h1>

      <button type="submit" class="action">
        <i class="material-icons">save</i>
      </button>

      <button class="action red">
        <i class="material-icons">delete</i>
      </button>

      <div class="card-action">
        <button
          v-if="!isNew"
          @click.prevent="deletePrompt"
          type="button"
          class="button button--flat button--red"
          :aria-label="$t('buttons.delete')"
          :title="$t('buttons.delete')">{{ $t('buttons.delete') }}</button>
        <input
          class="button button--flat"
          type="submit"
          :value="$t('buttons.save')">
      </div>
    </header>

    <main>
      <fieldset>
        <user-form :user.sync="user" :isDefault="false" :isNew="isNew" />
      </fieldset>
    </main>
  </form>

  <div v-if="$store.state.show === 'deleteUser'" class="card floating">
    <div class="card-content">
      <p>Are you sure you want to delete this user?</p>
    </div>

    <div class="card-action">
      <button class="button button--flat button--grey"
        @click="closeHovers"
        v-focus
        :aria-label="$t('buttons.cancel')"
        :title="$t('buttons.cancel')">
        {{ $t('buttons.cancel') }}
      </button>
      <button class="button button--flat"
        @click="deleteUser">
        {{ $t('buttons.delete') }}
      </button>
    </div>
  </div>
  </article>
</template>

<script>
import { mapMutations } from 'vuex'
import { users as api, settings } from '@/api'
import UserForm from '@/components/settings/UserForm'
import deepClone from 'lodash.clonedeep'

export default {
  name: 'user',
  components: {
    UserForm
  },
  data: () => {
    return {
      originalUser: null,
      user: {},
      loaded: false
    }
  },
  created () {
    this.fetchData()
  },
  computed: {
    isNew () {
      return this.$route.path === '/users/new'
    }
  },
  watch: {
    '$route': 'fetchData',
    'user.perm.admin': function () {
      if (!this.user.perm.admin) return
      this.user.lockPassword = false
    }
  },
  methods: {
    ...mapMutations([ 'closeHovers', 'showHover', 'setUser' ]),
    async fetchData () {
      try {
        if (this.isNew) {
          let { defaults } = await settings.get()
          this.user = {
            ...defaults,
            username: '',
            passsword: '',
            rules: [],
            drives: [{ icon: 'home', name: 'Home', path: '.' }],
            lockPassword: false,
            id: 0
          }
        } else {
          const id = this.$route.params.pathMatch
          this.user = { ...await api.get(id) }
        }

        this.loaded = true
      } catch (e) {
        this.$router.push({ path: '/settings/users/new' })
      }
    },
    deletePrompt () {
      this.showHover('deleteUser')
    },
    async deleteUser (event) {
      event.preventDefault()

      try {
        await api.remove(this.user.id)
        this.$router.push({ path: '/settings/users' })
        this.$showSuccess(this.$t('settings.userDeleted'))
      } catch (e) {
         (e.message === "403") ? this.$showError(this.$t("errors.forbidden"), false) : this.$showError(e)
      }
    },
    async save (event) {
      event.preventDefault()
      let user = {
        ...this.originalUser,
        ...this.user
      }

      try {
        if (this.isNew) {
          const loc = await api.create(user)
          this.$router.push({ path: loc })
          this.$showSuccess(this.$t('settings.userCreated'))
        } else {
          await api.update(user)

          if (user.id === this.$store.state.user.id) {
            this.setUser({ ...deepClone(user) })
          }

          this.$showSuccess(this.$t('settings.userUpdated'))
        }
      } catch (e) {
        this.$showError(e)
      }
    }
  }
}
</script>
