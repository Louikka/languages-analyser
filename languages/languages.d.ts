export interface LanguageDefinition {
    /** Name of the language. */
    name: string;
    /** List of shell filename/extension patterns used by the language. */
    matches: Array<string>;
}
