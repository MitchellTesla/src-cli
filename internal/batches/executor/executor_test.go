	"github.com/sourcegraph/src-cli/internal/batches/log"
	const rootPath = ""
	type filesByPath map[string][]string
	type filesByRepository map[string]filesByPath
				srcCLIRepo.ID: filesByPath{
					rootPath: []string{"README.md", "main.go"},
					rootPath: []string{"README.md"},
			name: "empty",
				{Run: "true"},
			// No diff should be generated.
				srcCLIRepo.ID: filesByPath{
					rootPath: []string{},
			name: "timeout",
				{Repo: srcCLIRepo, Files: map[string]string{"README.md": "line 1"}},
				// This needs to be a loop, because when a process goes to sleep
				// it's not interruptible, meaning that while it will receive SIGKILL
				// it won't exit until it had its full night of sleep.
				// So.
				// Instead we take short powernaps.
				{Run: `while true; do echo "zZzzZ" && sleep 0.05; done`},
			executorTimeout: 100 * time.Millisecond,
			wantErrInclude:  "execution in github.com/sourcegraph/src-cli failed: Timeout reached. Execution took longer than 100ms.",
			name: "templated steps",
				{Run: `go fmt main.go`},
				{Run: `touch modified-${{ join previous_step.modified_files " " }}.md`},
				{Run: `touch added-${{ join previous_step.added_files " " }}`},
					Run: `echo "hello.txt"`,
						"myOutput": batches.Output{
							Value: "${{ step.stdout }}",
				{Run: `touch output-${{ outputs.myOutput }}`},

				{Repository: srcCLIRepo},
				srcCLIRepo.ID: filesByPath{
					rootPath: []string{
						"main.go",
						"modified-main.go.md",
						"added-modified-main.go.md",
						"output-hello.txt",
					},
				{Repository: srcCLIRepo, Path: ""},
				{Repository: srcCLIRepo, Path: "a"},
				{Repository: srcCLIRepo, Path: "a/b"},
				srcCLIRepo.ID: filesByPath{
					rootPath: []string{"hello.txt", "gitignore-exists"},
					"a":      []string{"a/hello.txt", "a/gitignore-exists"},
					"a/b":    []string{"a/b/hello.txt", "a/b/gitignore-exists", "a/b/gitignore-exists-in-a"},
				srcCLIRepo.ID: filesByPath{
					rootPath: []string{"README.md"},
					"sub/directory/of/repo": []string{"README.md", "hello.txt", "in-path.txt"},
				},
			},
		},
		{
			name: "skips errors",
			archives: []mock.RepoArchive{
				{Repo: srcCLIRepo, Files: map[string]string{
					"README.md": "# Welcome to the README\n",
				}},
				{Repo: sourcegraphRepo, Files: map[string]string{
					"README.md": "# Sourcegraph README\n",
				}},
			},
			steps: []batches.Step{
				{Run: `echo -e "foobar\n" >> README.md`},
				{
					Run: `exit 1`,
					If:  fmt.Sprintf(`${{ eq repository.name %q }}`, sourcegraphRepo.Name),
				},
			},
			tasks: []*Task{
				{Repository: srcCLIRepo},
				{Repository: sourcegraphRepo},
			},
			wantFilesChanged: filesByRepository{
				srcCLIRepo.ID: filesByPath{
					rootPath: []string{"README.md"},
				sourcegraphRepo.ID: {},
			wantErrInclude: "execution in github.com/sourcegraph/sourcegraph failed: run: exit 1",
			// Make sure that the steps and tasks are setup properly
			for i := range tc.steps {
				tc.steps[i].SetImage(&mock.Image{
					RawDigest: tc.steps[i].Container,
				})
			}

			for _, task := range tc.tasks {
				task.BatchChangeAttributes = defaultBatchChangeAttributes
				task.Steps = tc.steps
			}

			// Setup a mock test server so we also test the downloading of archives

			// Setup an api.Client that points to this test server
			// Temp dir for log files and downloaded archives
			// Setup executor
			opts := newExecutorOpts{
				Creator: workspace.NewCreator(context.Background(), "bind", testTempDir, testTempDir, []batches.Step{}),
				Fetcher: batches.NewRepoFetcher(client, testTempDir, false),
				Logger:  log.NewManager(testTempDir, false),

			executor := newExecutor(opts)
			statusHandler := NewTaskStatusCollection([]*Task{})
			// Run executor
			executor.Start(context.Background(), tc.tasks, statusHandler)
			results, err := executor.Wait(context.Background())
			if tc.wantErrInclude == "" {
				if err != nil {
					t.Fatalf("execution failed: %s", err)
			} else {
				if err == nil {
					t.Fatalf("expected error to include %q, but got no error", tc.wantErrInclude)
						if !strings.Contains(strings.ToLower(err.Error()), strings.ToLower(tc.wantErrInclude)) {
			}
			wantResults := 0
			resultsFound := map[string]map[string]bool{}
			for repo, byPath := range tc.wantFilesChanged {
				wantResults += len(byPath)
				resultsFound[repo] = map[string]bool{}
				for path := range byPath {
					resultsFound[repo][path] = false
			}
			if have, want := len(results), wantResults; have != want {
				t.Fatalf("wrong number of execution results. want=%d, have=%d", want, have)
			}
			for _, taskResult := range results {
				repoID := taskResult.task.Repository.ID
				path := taskResult.task.Path
				wantFiles, ok := tc.wantFilesChanged[repoID]
				if !ok {
					t.Fatalf("unexpected file changes in repo %s", repoID)
				}
				resultsFound[repoID][path] = true
				wantFilesInPath, ok := wantFiles[path]
				if !ok {
					t.Fatalf("spec for repo %q and path %q but no files expected in that branch", repoID, path)
				fileDiffs, err := diff.ParseMultiFileDiff([]byte(taskResult.result.Diff))
				if err != nil {
					t.Fatalf("failed to parse diff: %s", err)
				if have, want := len(fileDiffs), len(wantFilesInPath); have != want {
					t.Fatalf("repo %s: wrong number of fileDiffs. want=%d, have=%d", repoID, want, have)
				diffsByName := map[string]*diff.FileDiff{}
				for _, fd := range fileDiffs {
					if fd.NewName == "/dev/null" {
						diffsByName[fd.OrigName] = fd
					} else {
						diffsByName[fd.NewName] = fd
					}
				for _, file := range wantFilesInPath {
					if _, ok := diffsByName[file]; !ok {
						t.Errorf("%s was not changed (diffsByName=%#v)", file, diffsByName)
			for repo, paths := range resultsFound {
				for path, found := range paths {
					for !found {
						t.Fatalf("expected spec to be created in path %s of repo %s, but was not", path, repo)
					}
				}