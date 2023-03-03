package main

import (
	"context"

	"github.com/google/go-github/v50/github"
	"github.com/samber/lo"
	"github.com/trim21/errgo"

	"github-bot/ent"
)

func (h PRHandle) setCheckRunFail(ctx context.Context, pull *ent.Pulls) error {
	_, _, err := h.g.Checks.UpdateCheckRun(ctx, pull.Owner, pull.Repo, pull.CheckRunID, github.UpdateCheckRunOptions{
		Name:       githubCheckRunName,
		Conclusion: &checkRunActionRequired,
		Output: &github.CheckRunOutput{
			Title:   lo.ToPtr("请关联你的 Bangumi 账号"),
			Summary: &checkRunDetailsMessage,
		},
	})
	return errgo.Trace(err)
}

func (h PRHandle) createCheckRun(ctx context.Context, owner, repo, headSha, result string) (*github.CheckRun, error) {
	r, _, err := h.g.Checks.CreateCheckRun(ctx, owner, repo, github.CreateCheckRunOptions{
		Name:       githubCheckRunName,
		HeadSHA:    headSha,
		Conclusion: &result,
		Output: &github.CheckRunOutput{
			Title:   lo.ToPtr("请关联你的 Bangumi 账号"),
			Summary: &checkRunDetailsMessage,
		},
	})
	if err != nil {
		return nil, errgo.Trace(err)
	}

	return r, nil
}

func (h PRHandle) createFailedCheckRun(ctx context.Context, owner, repo, headSha string) (*github.CheckRun, error) {
	r, _, err := h.g.Checks.CreateCheckRun(ctx, owner, repo, github.CreateCheckRunOptions{
		Name:       githubCheckRunName,
		HeadSHA:    headSha,
		Conclusion: &checkRunActionRequired,
		Output: &github.CheckRunOutput{
			Title:   lo.ToPtr("请关联你的 Bangumi 账号"),
			Summary: &checkRunDetailsMessage,
		},
	})
	if err != nil {
		return nil, errgo.Trace(err)
	}

	return r, nil
}
