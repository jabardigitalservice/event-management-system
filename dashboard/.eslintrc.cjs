module.exports = {
  root: true,
  env: {
    browser: true,
    node: true,
  },
  extends: [
    'plugin:@typescript-eslint/recommended',
    'plugin:nuxt/recommended',
    'plugin:vue/vue3-recommended',
    'prettier',
  ],
  parserOptions: {
    ecmaVersion: '2020',
    parser: '@typescript-eslint/parser',
    sourceType: 'module',
  },
  plugins: ['@typescript-eslint', 'vue', 'prettier'],
  rules: {
    'vue/script-setup-no-uses-vars': 'off',
    'vue/multi-word-component-names': [
      'error',
      {
        ignores: ['index', 'default'],
      },
    ],
    'vue/max-attributes-per-line': ["error", {
      'singleline': {
        'max': 3
      },
      'multiline': {
        'max': 1,
      }
    }],
  },
}
