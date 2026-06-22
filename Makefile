GREMLINS ?= go tool gremlins

# gremlins の全 mutator フラグ名（mappings.go / unleash の --help と一致）
MUTATORS := arithmetic-base conditionals-boundary conditionals-negation \
            increment-decrement invert-negatives invert-assignments \
            invert-bitwise invert-bwassign invert-logical invert-loopctrl \
            remove-self-assignments

# $(call only,<flag>) -> 指定した mutator だけを true、残りを false にするフラグ列
only = $(foreach m,$(MUTATORS),--$(m)=$(if $(filter $(m),$(1)),true,false))

.PHONY: test all \
        arithmetic-base conditionals-boundary conditionals-negation \
        increment-decrement invert-negatives invert-assignments \
        invert-bitwise invert-bwassign invert-logical invert-loopctrl \
        remove-self-assignments \
        status-runnable status-notcovered status-killed status-lived \
        status-notviable status-timedout status-skipped

# 元コードのテストが全てパスすることを確認（gremlins 実行の前提）
test:
	go test ./...

# リポジトリ全体に全 mutator を適用（.gremlins.yaml を使用）
all:
	$(GREMLINS) unleash ./...

# ---- mutator 別ターゲット（対象 mutator だけ有効化して単一変異に絞る） ----
arithmetic-base:
	$(GREMLINS) unleash $(call only,arithmetic-base) ./mutators/arithmetic_base
conditionals-boundary:
	$(GREMLINS) unleash $(call only,conditionals-boundary) ./mutators/conditionals_boundary
conditionals-negation:
	$(GREMLINS) unleash $(call only,conditionals-negation) ./mutators/conditionals_negation
increment-decrement:
	$(GREMLINS) unleash $(call only,increment-decrement) ./mutators/increment_decrement
invert-negatives:
	$(GREMLINS) unleash $(call only,invert-negatives) ./mutators/invert_negatives
invert-assignments:
	$(GREMLINS) unleash $(call only,invert-assignments) ./mutators/invert_assignments
invert-bitwise:
	$(GREMLINS) unleash $(call only,invert-bitwise) ./mutators/invert_bitwise
invert-bwassign:
	$(GREMLINS) unleash $(call only,invert-bwassign) ./mutators/invert_bwassign
invert-logical:
	$(GREMLINS) unleash $(call only,invert-logical) ./mutators/invert_logical
invert-loopctrl:
	$(GREMLINS) unleash $(call only,invert-loopctrl) ./mutators/invert_loopctrl
remove-self-assignments:
	$(GREMLINS) unleash $(call only,remove-self-assignments) ./mutators/remove_self_assignments

# ---- status 別ターゲット（-S で該当ステータスだけ表示） ----
# RUNNABLE: --dry-run でカバー済み変異を実行せず RUNNABLE のまま表示
status-runnable:
	$(GREMLINS) unleash -d -S r $(call only,arithmetic-base) ./mutators/arithmetic_base
# NOT COVERED: テスト未到達の行
status-notcovered:
	$(GREMLINS) unleash -S c $(call only,arithmetic-base) ./statuses/notcovered
# KILLED: テストが変異を検出
status-killed:
	$(GREMLINS) unleash -S k $(call only,arithmetic-base) ./mutators/arithmetic_base
# LIVED: 弱いテストが変異を見逃す
status-lived:
	$(GREMLINS) unleash -S l $(call only,arithmetic-base) ./statuses/lived
# NOT VIABLE: 変異後コードがコンパイル不能
status-notviable:
	$(GREMLINS) unleash -S v $(call only,arithmetic-base) ./statuses/notviable
# TIMED OUT: 変異が無限ループを生む
status-timedout:
	$(GREMLINS) unleash -S t $(call only,increment-decrement) ./statuses/timedout
# SKIPPED: --diff で差分外（作業ツリーがクリーンなら全て差分外）
status-skipped:
	$(GREMLINS) unleash -D HEAD -S s $(call only,arithmetic-base) ./mutators/arithmetic_base
