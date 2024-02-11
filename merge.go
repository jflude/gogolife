package main

type mergeable struct {
	eff []effect
	idx int
}

func newMergeable(eff []effect) *mergeable {
	return &mergeable{eff, 0}
}

func isUnmerged(m0 *mergeable) bool {
	return m0.idx < len(m0.eff)
}

func isEffectBefore(m0, m1 *mergeable) bool {
	return m0.eff[m0.idx].x < m1.eff[m1.idx].x
}

func merge1Effect(m0 *mergeable) (eff effect) {
	eff = m0.eff[m0.idx]
	m0.idx++
	return
}

func combine2Effects(m0, m1 *mergeable) effect {
	a0, a1 := merge1Effect(m0), merge1Effect(m1)
	return effect{a0.x, a0.delta + a1.delta}
}

func merge2Effects(m0, m1 *mergeable) effect {
	if isEffectBefore(m0, m1) {
		return merge1Effect(m0)
	} else if isEffectBefore(m1, m0) {
		return merge1Effect(m1)
	} else {
		return combine2Effects(m0, m1)
	}
}

func combine3Effects(m0, m1, m2 *mergeable) effect {
	a0, a1, a2 := merge1Effect(m0), merge1Effect(m1), merge1Effect(m2)
	return effect{a0.x, a0.delta + a1.delta + a2.delta}
}

func merge3Effects(m0, m1, m2 *mergeable) (eff effect, done bool) {
	if isUnmerged(m0) {
		if isUnmerged(m1) {
			if isUnmerged(m2) {
				if isEffectBefore(m0, m1) {
					eff = merge2Effects(m0, m2)
				} else if isEffectBefore(m1, m0) {
					eff = merge2Effects(m1, m2)
				} else if isEffectBefore(m2, m0) {
					eff = merge1Effect(m2)
				} else if isEffectBefore(m0, m2) {
					eff = combine2Effects(m0, m1)
				} else {
					eff = combine3Effects(m0, m1, m2)
				}
			} else {
				eff = merge2Effects(m0, m1)
			}
		} else if isUnmerged(m2) {
			eff = merge2Effects(m0, m2)
		} else {
			eff = merge1Effect(m0)
		}
	} else if isUnmerged(m1) {
		if isUnmerged(m2) {
			eff = merge2Effects(m1, m2)
		} else {
			eff = merge1Effect(m1)
		}
	} else if isUnmerged(m2) {
		eff = merge1Effect(m2)
	} else {
		done = true
	}

	return
}
