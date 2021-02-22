<template>
  <article id="settings" v-if="settings !== null">
  <form @submit.prevent="save">
    <header id="title">
      <h1>{{ $t('settings.globalSettings') }}</h1>

      <button type="submit" class="action">
        <i class="material-icons">save</i>
      </button>
    </header>

    <main>
      <fieldset>
        <h2>{{ $t('settings.globalSettings') }}</h2>

        <label>
          {{ $t('settings.instanceName') }}
          <input class="input input--block" type="text" v-model="settings.branding.name">
        </label>

        <label>
          {{ $t('settings.rootFolder') }}
          <input class="input input--block" type="text">
        </label>
        <!--
        <label for="branding-files">
          {{ $t('settings.brandingDirectoryPath') }}
          <input class="input input--block" type="text" v-model="settings.branding.files" id="branding-files" />
        </label>
        -->

        <h3>{{ $t('settings.userManagement') }}</h3>
        <label><input type="checkbox" v-model="settings.signup"> {{ $t('settings.allowSignup') }}</label>
        <label><input type="checkbox" v-model="settings.createUserDir"> {{ $t('settings.createUserDir') }}</label>
      </fieldset>

      <fieldset>
        <h2>{{ $t('settings.userDefaults') }}</h2>
        <p class="small">{{ $t('settings.defaultUserDescription') }}</p>

          <user-form :isNew="false" :isDefault="true" :user.sync="settings.defaults" />
      </fieldset>
    </main>
<!--
    <form v-if="isExecEnabled" @submit.prevent="save">
      <h2>{{ $t('settings.commandRunner') }}</h2>
      <button type="submit" class="action">
        <i class="material-icons">check</i>
      </button>

      <div class="content">
        <i18n path="settings.commandRunnerHelp" tag="p" class="small">
          <code>FILE</code>
          <code>SCOPE</code>
          <a class="link" target="_blank" href="https://filebrowser.org/configuration/command-runner">{{ $t('settings.documentation') }}</a>
        </i18n>

        <div v-for="command in settings.commands" :key="command.name" class="collapsible">
          <input :id="command.name" type="checkbox">
          <label :for="command.name">
            <p>{{ capitalize(command.name) }}</p>
            <i class="material-icons">arrow_drop_down</i>
          </label>
          <div class="collapse">
            <textarea class="input input--block input--textarea" v-model.trim="command.value"></textarea>
          </div>
        </div>
      </div>
    </form>
-->
  </form>
  </article>
</template>

<script>
import { mapState } from 'vuex'
import { settings as api } from '@/api'
import UserForm from '@/components/settings/UserForm'
import { enableExec } from '@/utils/constants'

export default {
  name: 'settings',
  components: {
    UserForm
  },
  data: function () {
    return {
      originalSettings: null,
      settings: null
    }
  },
  computed: {
    ...mapState([ 'user' ]),
    isExecEnabled: () => enableExec
  },
  async created () {
    try {
      const original = await api.get()
      let settings = { ...original, commands: [] }

      for (const key in original.commands) {
        settings.commands.push({
          name: key,
          value: original.commands[key].join('\n')
        })
      }

      this.originalSettings = original
      this.settings = settings
    } catch (e) {
      this.$showError(e)
    }
  },
  methods: {
    capitalize (name, where = '_') {
      if (where === 'caps') where = /(?=[A-Z])/
      let splitted = name.split(where)
      name = ''

      for (let i = 0; i < splitted.length; i++) {
        name += splitted[i].charAt(0).toUpperCase() + splitted[i].slice(1) + ' '
      }

      return name.slice(0, -1)
    },
    async save () {
      let settings = {
        ...this.settings,
        commands: {}
      }

      for (const { name, value } of this.settings.commands) {
        settings.commands[name] = value.split('\n').filter(cmd => cmd !== '')
      }

      try {
        await api.update(settings)
        this.$showSuccess(this.$t('settings.settingsUpdated'))
      } catch (e) {
        this.$showError(e)
      }
    }
  }
}
</script>
