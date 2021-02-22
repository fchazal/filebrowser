<template>
  <div>
    <h3>{{ $t('settings.permissions') }}</h3>
    <p class="small">{{ $t('settings.permissionsHelp') }}</p>

    <label>
      <input type="checkbox" v-model="admin">
      {{ $t('settings.administrator') }}
    </label>

    <label>
      <input type="checkbox" :disabled="admin" v-model="perm.create">
      {{ $t('settings.perm.create') }}
    </label>

    <label>
      <input type="checkbox" :disabled="admin" v-model="perm.delete">
      {{ $t('settings.perm.delete') }}
    </label>

    <label>
      <input type="checkbox" :disabled="admin" v-model="perm.download">
      {{ $t('settings.perm.download') }}
    </label>

    <label>
      <input type="checkbox" :disabled="admin" v-model="perm.modify">
      {{ $t('settings.perm.modify') }}
    </label>

    <label>
      <input type="checkbox" :disabled="admin" v-model="perm.rename">
      {{ $t('settings.perm.rename') }}
    </label>

    <label>
      <input type="checkbox" :disabled="admin" v-model="perm.share">
      {{ $t('settings.perm.share') }}
    </label>
  </div>
</template>

<script>
import { enableExec } from '@/utils/constants'
export default {
  name: 'permissions',
  props: ['perm'],
  computed: {
    admin: {
      get () {
        return this.perm.admin
      },
      set (value) {
        if (value) {
          for (const key in this.perm) {
            this.perm[key] = true
          }
        }

        this.perm.admin = value
      }
    },
    isExecEnabled: () => enableExec
  }
}
</script>
