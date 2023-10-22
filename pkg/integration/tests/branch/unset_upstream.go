package branch

import (
	"github.com/jesseduffield/lazygit/pkg/config"
	. "github.com/jesseduffield/lazygit/pkg/integration/components"
)

var UnsetUpstream = NewIntegrationTest(NewIntegrationTestArgs{
	Description:  "fsjalll",
	ExtraCmdArgs: []string{},
	Skip:         false,
	SetupConfig:  func(config *config.AppConfig) {},
	SetupRepo: func(shell *Shell) {
		shell.
			EmptyCommit("one").
			CloneIntoRemote("origin").
			NewBranch("test_branch").
			PushBranch("origin", "test_branch").
			Checkout("master").
			SetBranchUpstream("master", "origin/master").
			RemoveRemoteBranch("origin", "test_branch")
	},
	Run: func(t *TestDriver, keys config.KeybindingConfig) {
		t.Views().Branches().
			Focus().
			Press(keys.Universal.NextScreenMode). // we need to enlargen the window to see the upstream
			SelectedLines(
				Contains("master").Contains("origin master"),
			).
			Press(keys.Branches.SetUpstream).
			Tap(func() {
				t.ExpectPopup().Menu().
					Title(Equals("Upstream options")).
					Select(Contains("Unset upstream of selected branch")).
					Confirm()
			}).
			SelectedLines(
				Contains("master").DoesNotContain("origin master"),
			)

		t.Views().Branches().
			Focus().
			SelectNextItem().
			SelectedLines(
				Contains("test_branch").Contains("upstream gone"),
			).
			Press(keys.Branches.SetUpstream).
			Tap(func() {
				t.ExpectPopup().Menu().
					Title(Equals("Upstream options")).
					Select(Contains("Unset upstream of selected branch")).
					Confirm()
			}).
			SelectedLines(
				Contains("test_branch").DoesNotContain("origin test_branch").DoesNotContain("upstream gone"),
			)
	},
})
