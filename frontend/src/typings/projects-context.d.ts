interface ProjectsContext {
  projects?: WorkQuery['projects'];
  loading: boolean;
  activeFilters: Array<string>;
  setActiveFilter: (name: string) => void;
  clearFilters: () => void;
}
