;;; nupload --- Upload files to nup.pw
;; Copyright (C) 2015 pinku

;; Author: pinku <pinku@pinku.pw>
;; Created: 12 Jul 2015
;; Version: 0.0.1

;; This file is not part of GNU Emacs.

;; This file is free software: you can redistribute it and/or modify
;; it under the terms of the GNU General Public License as published by
;; the Free Software Foundation, either version 3 of the License, or
;; (at your option) any later version.
;; This file is distributed in the hope that it will be useful,
;; but WITHOUT ANY WARRANTY; without even the implied warranty of
;; MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
;; GNU General Public License for more details.
;; You should have received a copy of the GNU General Public License
;; along with this file.  If not, see <http://www.gnu.org/licenses/>.

;;; Commentary:

;;; Code:

;; --- constants
(defconst nupload-upload-url "https://nup.pw")
(defconst nupload-hyperboria-upload-url "http://h.nup.pw")
(defconst nupload--regex "<a href=\"\\(https?://\\(?:h\.\\)?nup\.pw/.+?\\)\"")

;; ---- config options
(defcustom nupload-use-hyperboria nil
  "Connect to nup.pw via Hyperboria"
  :type 'boolean)

(defun nupload--url ()
  (if nupload-use-hyperboria nupload-hyperboria-upload-url nupload-upload-url))

(defun nupload--parse (s)
  (string-match nupload--regex s)
  (match-string 1 s))

(defun nupload-buffer ()
  (interactive)
  (let ((f (make-temp-file "nupload"))
	(c (current-buffer)))
    (with-temp-buffer
      (insert-buffer c)
      (write-file f)
      (nupload-file f))
    (delete-file f)))

(defun nupload-file (filename)
  (interactive (list (expand-file-name (read-file-name "Filename: "))))
  (message (nupload--parse (shell-command-to-string (format "curl -sF \"file=@%s\" %s" filename (nupload--url))))))

(defun nupload ()
  (interactive)
  (cond
   ((string= major-mode "dired-mode")
    (nupload-file (dired-get-filename)))
   ((buffer-file-name) (nupload-file (buffer-file-name)))
   (t (nupload-buffer))))

(provide 'nupload)
