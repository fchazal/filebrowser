<template>
  <form class="rules">
    <div v-for="(rule, index) in rules" :key="index">
      <input
        @keypress.enter.prevent
        type="text"
        v-if="rule.regex"
        v-model="rule.regexp.raw"
        :placeholder="$t('settings.insertRegex')" />

      <button class="action red" @click="remove($event, index)">
        <i class="material-icons">delete</i>
      </button>
    </div>

    <button class="action long" @click="create" default="false">
      <i class="material-icons">add</i>
      {{ $t('buttons.addRule') }}
    </button>
  </form>
</template>

<script>
export default {
  name: 'rules-textarea',
  props: ['rules'],
  methods: {
    remove (event, index) {
      event.preventDefault()
      let rules = [ ...this.rules ]
      rules.splice(index, 1)
      this.$emit('update:rules', [ ...rules ])
    },
    create (event) {
      event.preventDefault()

      this.$emit('update:rules', [
        ...this.rules,
        {
          allow: false,
          path: '',
          regex: true,
          regexp: {
            raw: ''
          }
        }
      ])
    }
  }
}
</script>
