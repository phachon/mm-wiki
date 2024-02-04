module.exports = {
  parser: '@typescript-eslint/parser',
  plugins: ['react', 'react-hooks', '@typescript-eslint/eslint-plugin', 'prettier'],
  globals: {
    definePageConfig: true,
    requirePlugin: true
  },
  env: {
    node: true, // 只需将该项设置为 true 即可
    browser: true,
    es6: true
  },
  rules: {
    // 这里自定义规则，规则地址:
    // http://eslint.cn/docs/rules/
    'no-dupe-keys': 2, //在创建对象字面量时不允许键重复 {a:1,a:1}
    'no-dupe-args': 2, //函数参数不能重复
    'no-extra-semi': 2, //禁止多余的冒号
    'no-func-assign': 2, //禁止重复的函数声明
    'no-irregular-whitespace': 2, //不能有不规则的空格
    'no-multi-spaces': 1, //不能用多余的空格
    'no-param-reassign': 0, //禁止给参数重新赋值
    'no-spaced-func': 2, //函数调用时 函数名与()之间不能有空格
    'no-trailing-spaces': 1, //一行结束后面不要有空格
    'no-undef': 1, //不能有未定义的变量
    'no-use-before-define': 'off', //未定义前不能使用
    'no-var': 0, //禁用var，用let和const代替
    'no-extra-parens': 0,
    'vue/no-unused-components': 'off',
    'array-bracket-spacing': [2, 'never'], //是否允许非空数组里面有多余的空格
    camelcase: 1, //强制驼峰法命名
    'comma-dangle': [2, 'never'], //对象字面量项尾不能有逗号
    quotes: [1, 'single'], //引号类型 `` "" ''
    semi: [2, 'never'], //语句分号结尾
    'semi-spacing': [
      0,
      {
        before: false,
        after: true
      }
    ], //分号前后空格
    'space-after-keywords': [0, 'always'], //关键字后面是否要空一格
    'space-before-blocks': [0, 'always'], //不以新行开始的块{前面要不要有空格
    'space-before-function-paren': [0, 'always'], //函数定义时括号前面要不要有空格
    'space-in-parens': [0, 'never'], //小括号里面要不要有空格
    'arrow-spacing': 0, //=>的前/后括号
    'linebreak-style': 0,
    allowForLoopAfterthoughts: 0,
    'key-spacing': 0,
    'prettier/prettier': 'error',
    'react/jsx-uses-react': 'off',
    'react/react-in-jsx-scope': 'off'
  }
}
