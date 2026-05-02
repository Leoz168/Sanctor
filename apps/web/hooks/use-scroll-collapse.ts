"use client";

import { useEffect, useState } from "react";

export function useScrollCollapse(topThreshold = 8) {
  const [isExpanded, setIsExpanded] = useState(true);

  useEffect(() => {
    const syncExpandedState = () => {
      setIsExpanded(window.scrollY <= topThreshold);
    };

    syncExpandedState();
    window.addEventListener("scroll", syncExpandedState, { passive: true });

    return () => {
      window.removeEventListener("scroll", syncExpandedState);
    };
  }, [topThreshold]);

  return isExpanded;
}
