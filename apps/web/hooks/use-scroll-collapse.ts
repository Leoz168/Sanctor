"use client";

import { useEffect, useRef, useState } from "react";

export function useScrollCollapse(expandDelay = 180) {
  const [isExpanded, setIsExpanded] = useState(true);
  const timeoutRef = useRef<ReturnType<typeof setTimeout> | null>(null);

  useEffect(() => {
    const handleScroll = () => {
      setIsExpanded(false);

      if (timeoutRef.current) {
        clearTimeout(timeoutRef.current);
      }

      timeoutRef.current = setTimeout(() => {
        setIsExpanded(true);
      }, expandDelay);
    };

    window.addEventListener("scroll", handleScroll, { passive: true });

    return () => {
      window.removeEventListener("scroll", handleScroll);

      if (timeoutRef.current) {
        clearTimeout(timeoutRef.current);
      }
    };
  }, [expandDelay]);

  return isExpanded;
}
