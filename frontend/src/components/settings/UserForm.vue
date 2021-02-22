<template>
  <div>
    <label v-if="!isDefault">
      {{ $t('settings.username') }}
      <input class="input input--block" type="text" v-model="user.username">
    </label>

    <label v-if="!isDefault">
      <input type="checkbox" :disabled="user.perm.admin" v-model="user.lockPassword">
      {{ $t('settings.lockPassword') }}
    </label>
    
    <label v-if="!isDefault" for="password">
      {{ $t('settings.password') }}
      <input class="input input--block" type="password" :placeholder="passwordPlaceholder" v-model="user.password" id="password">
    </label>

    <label for="scope">
      {{ $t('settings.scope') }}
      <input class="input input--block" type="text" v-model="user.scope" id="scope">
    </label>

    <label for="locale">
      {{ $t('settings.language') }}
      <languages class="input input--block" id="locale" :locale.sync="user.locale"></languages>
    </label>
  
    <label>
      <input type="checkbox" v-model="user.hideDotfiles">
      {{ $t('settings.hideDotfiles') }}
    </label>
    
    <label>
      <input type="checkbox" v-model="user.singleClick">
      {{ $t('settings.singleClick') }}
    </label>

    <permissions :perm.sync="user.perm" />

    <div v-if="!isDefault">
      <h3>{{ $t('settings.rules') }}</h3>
      <p class="small">{{ $t('settings.rulesHelp') }}</p>
      <rules :rules.sync="user.rules" />
    </div>
  </div>
</template>

<script>
import Languages from './Languages'
import Rules from './Rules'
import Permissions from './Permissions'
import { enableExec } from '@/utils/constants'

export default {
  name: 'user',
  components: {
    Permissions,
    Languages,
    Rules,
  },
  props: [ 'user', 'isNew', 'isDefault' ],
  computed: {
    passwordPlaceholder () {
      return this.isNew ? '' : this.$t('settings.avoidChanges')
    },
    isExecEnabled: () => enableExec
  },
  watch: {
    'user.perm.admin': function () {
      if (!this.user.perm.admin) return
      this.user.lockPassword = false
    }
  }
}
</script>
