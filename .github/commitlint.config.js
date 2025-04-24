module.exports = {
    extends: ['@commitlint/config-conventional'],
    rules: {
        'type-enum': [
            2,
            'always',
            [
                'feat',     // @note: New features
                'fix',      // @note: Bug fixes
                'docs',     // @note: Documentation changes
                'style',    // @note: Code style changes (formatting, semicolons, etc)
                'refactor', // @note: Code refactoring
                'perf',     // @note: Performance improvements
                'test',     // @note: Adding or updating tests
                'chore',    // @note: Maintenance tasks, upgrades, etc
                'ci',       // @note: CI/CD related changes
                'revert',   // @note: Reverting changes
                'build'     // @note: Build system or external dependencies
            ]
        ],
        'type-case': [2, 'always', 'lowercase'],
        'type-empty': [2, 'never'],
        'scope-case': [2, 'always', ['lower-case', 'kebab-case']],
        'subject-case': [
            2,
            'never',
            ['upper-case', 'pascal-case', 'camel-case']
        ],
        'subject-empty': [2, 'never'],
        'subject-full-stop': [2, 'never', '.'],
        'header-max-length': [2, 'always', 72],
        'body-leading-blank': [1, 'always'],
        'footer-leading-blank': [1, 'always'],
        'footer-max-line-length': [2, 'always', 100]
    }
};
