GREMLINS ?= go tool gremlins

# Put the Go build cache in a writable location outside the module so that
# `go test` / gremlins keep working even where the default GOCACHE is read-only
# (sandboxes etc.). Keeping it inside the module would pull it into the working
# directory copy that gremlins makes, so it must live outside (under $TMPDIR).
GOCACHE ?= $(or $(TMPDIR),/tmp)/gremlins-showcase-gocache
export GOCACHE

# The gremlins timeout = coverage measurement time x coefficient (default 3).
# With a warm cache, coverage measurement finishes almost instantly and a
# coefficient of 3 leaves no time to recompile the mutated code, causing a
# spurious TIMED OUT. Raise the coefficient to keep things stable.
TC := --timeout-coefficient 30

# All mutator flag names of gremlins (matching mappings.go / `unleash --help`).
MUTATORS := arithmetic-base conditionals-boundary conditionals-negation \
            increment-decrement invert-negatives invert-assignments \
            invert-bitwise invert-bwassign invert-logical invert-loopctrl \
            remove-self-assignments

# $(call only,<flag>) -> a flag sequence that sets only the given mutator to true and the rest to false.
only = $(foreach m,$(MUTATORS),--$(m)=$(if $(filter $(m),$(1)),true,false))

.PHONY: test all \
        arithmetic-base conditionals-boundary conditionals-negation \
        increment-decrement invert-negatives invert-assignments \
        invert-bitwise invert-bwassign invert-logical invert-loopctrl \
        remove-self-assignments \
        status-runnable status-notcovered status-killed status-lived \
        status-notviable status-timedout status-skipped

# Verify that all tests on the original code pass (a prerequisite for running gremlins).
test:
	go test ./...

# Apply all mutators to the whole repository (uses .gremlins.yaml).
all:
	$(GREMLINS) unleash $(TC) ./...

# ---- per-mutator targets (enable only the target mutator to isolate a single mutation) ----
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

# ---- per-status targets (use -S to show only the relevant status) ----
# RUNNABLE: --dry-run shows covered mutations as RUNNABLE without executing them.
status-runnable:
	$(GREMLINS) unleash $(TC) -d -S r $(call only,arithmetic-base) ./mutators/arithmetic_base
# NOT COVERED: a line not reached by any test.
status-notcovered:
	$(GREMLINS) unleash $(TC) -S c $(call only,arithmetic-base) ./statuses/notcovered
# KILLED: a test detects the mutation.
status-killed:
	$(GREMLINS) unleash $(TC) -S k $(call only,arithmetic-base) ./mutators/arithmetic_base
# LIVED: a weak test misses the mutation.
status-lived:
	$(GREMLINS) unleash $(TC) -S l $(call only,arithmetic-base) ./statuses/lived
# NOT VIABLE: the mutated code fails to compile.
status-notviable:
	$(GREMLINS) unleash $(TC) -S v $(call only,arithmetic-base) ./statuses/notviable
# TIMED OUT: the mutation produces an infinite loop.
status-timedout:
	$(GREMLINS) unleash $(TC) -S t $(call only,increment-decrement) ./statuses/timedout
# SKIPPED: in --diff mode, a mutation that falls "outside the changed diff" is reported as SKIPPED.
# The most recent commit (HEAD) does not modify arithmetic_base.go, so the diff against HEAD~1
# does not include arithmetic_base.go and its mutation is outside the diff = SKIPPED.
# Note: with an empty diff (clean working tree vs HEAD) gremlins disables the diff filter and
# runs everything (no SKIPPED). SKIPPED requires a non-empty diff.
# If a future commit that modifies arithmetic_base.go is stacked on the tip, its mutation falls
# inside the diff and is no longer SKIPPED. In that case, point -D at a ref before the
# arithmetic_base.go change.
status-skipped:
	$(GREMLINS) unleash $(TC) -D HEAD~1 -S s $(call only,arithmetic-base) ./mutators/arithmetic_base
