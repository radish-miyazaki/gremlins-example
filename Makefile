GREMLINS ?= go tool gremlins

# go のビルドキャッシュをモジュール外の書き込み可能な場所に置く。
# 既定の GOCACHE が書き込み不可な環境（サンドボックス等）でも go test / gremlins が
# 動くようにするため。モジュール内に置くと gremlins の作業ディレクトリ複製に含まれて
# しまうため、必ずモジュール外（$TMPDIR 配下）に置く。
GOCACHE ?= $(or $(TMPDIR),/tmp)/gremlins-showcase-gocache
export GOCACHE

# gremlins のタイムアウト = カバレッジ計測時間 × 係数（既定3）。
# キャッシュが温まっているとカバレッジ計測が一瞬で終わり、係数3では変異後の
# 再コンパイルが間に合わず誤って TIMED OUT になる。係数を上げて安定させる。
TC := --timeout-coefficient 30

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
	$(GREMLINS) unleash $(TC) ./...

# ---- mutator 別ターゲット（対象 mutator だけ有効化して単一変異に絞る） ----
arithmetic-base:
	$(GREMLINS) unleash $(TC) $(call only,arithmetic-base) ./mutators/arithmetic_base
conditionals-boundary:
	$(GREMLINS) unleash $(TC) $(call only,conditionals-boundary) ./mutators/conditionals_boundary
conditionals-negation:
	$(GREMLINS) unleash $(TC) $(call only,conditionals-negation) ./mutators/conditionals_negation
increment-decrement:
	$(GREMLINS) unleash $(TC) $(call only,increment-decrement) ./mutators/increment_decrement
invert-negatives:
	$(GREMLINS) unleash $(TC) $(call only,invert-negatives) ./mutators/invert_negatives
invert-assignments:
	$(GREMLINS) unleash $(TC) $(call only,invert-assignments) ./mutators/invert_assignments
invert-bitwise:
	$(GREMLINS) unleash $(TC) $(call only,invert-bitwise) ./mutators/invert_bitwise
invert-bwassign:
	$(GREMLINS) unleash $(TC) $(call only,invert-bwassign) ./mutators/invert_bwassign
invert-logical:
	$(GREMLINS) unleash $(TC) $(call only,invert-logical) ./mutators/invert_logical
invert-loopctrl:
	$(GREMLINS) unleash $(TC) $(call only,invert-loopctrl) ./mutators/invert_loopctrl
remove-self-assignments:
	$(GREMLINS) unleash $(TC) $(call only,remove-self-assignments) ./mutators/remove_self_assignments

# ---- status 別ターゲット（-S で該当ステータスだけ表示） ----
# RUNNABLE: --dry-run でカバー済み変異を実行せず RUNNABLE のまま表示
status-runnable:
	$(GREMLINS) unleash $(TC) -d -S r $(call only,arithmetic-base) ./mutators/arithmetic_base
# NOT COVERED: テスト未到達の行
status-notcovered:
	$(GREMLINS) unleash $(TC) -S c $(call only,arithmetic-base) ./statuses/notcovered
# KILLED: テストが変異を検出
status-killed:
	$(GREMLINS) unleash $(TC) -S k $(call only,arithmetic-base) ./mutators/arithmetic_base
# LIVED: 弱いテストが変異を見逃す
status-lived:
	$(GREMLINS) unleash $(TC) -S l $(call only,arithmetic-base) ./statuses/lived
# NOT VIABLE: 変異後コードがコンパイル不能
status-notviable:
	$(GREMLINS) unleash $(TC) -S v $(call only,arithmetic-base) ./statuses/notviable
# TIMED OUT: 変異が無限ループを生む
status-timedout:
	$(GREMLINS) unleash $(TC) -S t $(call only,increment-decrement) ./statuses/timedout
# SKIPPED: --diff モードで「変更された差分の外」にある変異は SKIPPED になる。
# 直近コミット(HEAD)は arithmetic_base.go を変更していないため、HEAD~1 との差分には
# arithmetic_base.go が含まれず、その変異は差分外 = SKIPPED となる。
# 注意: 空 diff（クリーンな作業ツリー vs HEAD）では gremlins は diff フィルタを無効化し
# 全て実行してしまう（SKIPPED にならない）。SKIPPED には「非空の差分」が必要。
# 将来 arithmetic_base.go を変更するコミットを tip に積むと、その変異が差分内に入り
# SKIPPED にならなくなる。その場合は -D を arithmetic_base.go 変更より前の ref に向けること。
status-skipped:
	$(GREMLINS) unleash $(TC) -D HEAD~1 -S s $(call only,arithmetic-base) ./mutators/arithmetic_base
