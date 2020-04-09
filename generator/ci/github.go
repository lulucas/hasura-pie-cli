package ci

import "github.com/lulucas/hasura-pie-cli/utils"

const githubCITpl = `name: api
on: [push]

jobs:
  business:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@master
      - uses: actions/setup-go@v1
        with:
          go-version: 1.14
      - uses: actions/cache@v1
        with:
          path: ~/go/pkg/mod
          key: ${{ runner.os }}-go-${{ hashFiles('**/go.sum') }}
          restore-keys: |
            ${{ runner.os }}-go-
      - env:
          CGO_ENABLED: 0
        run: |
          go build -o functions/business/main functions/business/main.go
      - env:
          SSH_HOST: ${{ secrets.BUSINESS_SSH_HOST }}
          SSH_USER: ${{ secrets.BUSINESS_SSH_USER }}
          SSH_PASS: ${{ secrets.BUSINESS_SSH_PASS }}
        run: |
          mkdir -p ~/.ssh && ssh-keyscan $SSH_HOST >> ~/.ssh/known_hosts
          cd functions/business
          sshpass -p $SSH_PASS rsync -avzr --include-from=.rsync --chmod=D755,F755 . $SSH_USER@$SSH_HOST:/app --delete
          sshpass -p $SSH_PASS ssh $SSH_USER@$SSH_HOST "cd /app && docker-compose down --remove-orphans && docker-compose up --build -d"
`

func GenerateGithubAction() error {
	if err := utils.SaveToFile(".github/workflows/ci.yml", githubCITpl); err != nil {
		return err
	}
	return nil
}
