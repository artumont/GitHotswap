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
        'type-case': [2, 'always', 'lower'],
        'type-empty': [2, 'never'],
        'subject-case': [2, 'always', 'sentence-case'],
        'subject-empty': [2, 'never'],
        'subject-full-stop': [2, 'never', '.'],
        'subject-min-length': [2, 'always', 5],
        'body-leading-blank': [2, 'always'],
        'header-max-length': [2, 'always', 72],
        'footer-leading-blank': [1, 'always']
    }
};
