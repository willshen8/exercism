export const abilityModifier = (score) => {
  if (score < 3) throw new Error("Ability scores must be at least 3");
  if (score > 18) throw new Error("Ability scores can be at most 18");
  return Math.floor((score - 10) / 2);
};

export class Character {
  constructor() {
    this.strength = Character.rollAbility();
    this.dexterity = Character.rollAbility();
    this.constitution = Character.rollAbility();
    this.intelligence = Character.rollAbility();
    this.wisdom = Character.rollAbility();
    this.charisma = Character.rollAbility();
  }

  static rollAbility() {
    var randomAbilityScores = [];
    for (let i = 0; i < 4; i++) {
      randomAbilityScores[i] = Math.floor(Math.random() * 6 + 1);
    }
    return randomAbilityScores
      .sort()
      .slice(1)
      .reduce((a, b) => a + b);
  }

  get hitpoints() {
    return 10 + abilityModifier(this.constitution);
  }
}
